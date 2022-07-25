#![allow(unused_imports)]
#![allow(clippy::single_component_path_imports)]
//#![feature(backtrace)]

use std::io::{Read, Write};
use std::net::{TcpListener, TcpStream};
use std::path::PathBuf;
use std::sync::{Condvar, Mutex};
use std::{cell::RefCell, env, sync::atomic::*, sync::Arc, thread, time::*};
use std::{fs, time};

use anyhow::bail;

use embedded_svc::mqtt::client::utils::ConnState;
use log::*;

use url;

use smol;

use embedded_hal::adc::OneShot;
use embedded_hal::blocking::delay::DelayMs;
use embedded_hal::digital::v2::OutputPin;

use embedded_svc::eth;
use embedded_svc::eth::{Eth, TransitionalState};
use embedded_svc::httpd::registry::*;
use embedded_svc::httpd::*;
use embedded_svc::io;
use embedded_svc::ipv4;
use embedded_svc::mqtt::client::{Client, Connection, MessageImpl, Publish, QoS};
use embedded_svc::ping::Ping;
use embedded_svc::sys_time::SystemTime;
use embedded_svc::timer::TimerService;
use embedded_svc::timer::*;
use embedded_svc::wifi::*;

use esp_idf_svc::eth::*;
use esp_idf_svc::eventloop::*;
use esp_idf_svc::eventloop::*;
use esp_idf_svc::httpd as idf;
use esp_idf_svc::httpd::ServerRegistry;
use esp_idf_svc::mqtt::client::*;
use esp_idf_svc::netif::*;
use esp_idf_svc::nvs::*;
use esp_idf_svc::ping;
use esp_idf_svc::sntp;
use esp_idf_svc::sysloop::*;
use esp_idf_svc::systime::EspSystemTime;
use esp_idf_svc::timer::*;
use esp_idf_svc::wifi::*;

use esp_idf_hal::adc;
use esp_idf_hal::delay;
use esp_idf_hal::gpio;
use esp_idf_hal::i2c;
use esp_idf_hal::prelude::*;
use esp_idf_hal::spi;

use esp_idf_sys::{self, c_types};
use esp_idf_sys::{esp, EspError};

use display_interface_spi::SPIInterfaceNoCS;

use embedded_graphics::mono_font::{ascii::FONT_10X20, MonoTextStyle};
use embedded_graphics::pixelcolor::*;
use embedded_graphics::prelude::*;
use embedded_graphics::primitives::*;
use embedded_graphics::text::*;

use ili9341;
use ssd1306;
use ssd1306::mode::DisplayConfig;
use st7789;

use ads1x1x::{
    channel, Ads1x1x, ComparatorLatching, ComparatorMode, ComparatorPolarity, ComparatorQueue,
    DataRate12Bit, FullScaleRange, ModeChangeError, SlaveAddr,
};
use bytes::{BufMut, BytesMut};
use nb::block;

use hex;

fn main() -> Result<()> {
    // breker conf
    let device_token = base64::decode_config(
        "YE6xYc-GG5cFOwsU7pPSx2M_SGeiLkE_jJbSGWLd0UzPhhtCV3du8WLd0UzPhhtCV3du8mLd0UzPhhtCV3du8w",
        base64::URL_SAFE_NO_PAD,
    )
    .unwrap();

    esp_idf_sys::link_patches();

    // Bind the log crate to the ESP Logging facilities
    esp_idf_svc::log::EspLogger::initialize_default();

    let peripherals = Peripherals::take().unwrap();
    let pins = peripherals.pins;

    let i2c = peripherals.i2c0;
    let scl = pins.gpio22;
    let sda = pins.gpio21;

    let config = <i2c::config::MasterConfig as Default>::default().baudrate(9600.Hz().into());
    let dev = i2c::Master::<i2c::I2C0, _, _>::new(i2c, i2c::MasterPins { sda, scl }, config)?;

    let address = SlaveAddr::default();
    let mut adc = Ads1x1x::new_ads1015(dev, address);
    adc.set_data_rate(DataRate12Bit::Sps3300).unwrap();
    adc.set_full_scale_range(FullScaleRange::Within1_024V)
        .unwrap();

    match adc.into_continuous() {
        Err(ModeChangeError::I2C(_e, _adc)) =>
        /* mode change failed handling */
        {
            panic!()
        }
        Ok(mut adc) => {
            let mut readings = Vec::with_capacity(1000);

            let mut second_timer = time::Instant::now();
            // let mut batch_timer = time::Instant::now();

            loop {
                let measurement = adc.read().unwrap();
                let a = volts_to_amps(measure_to_volts(measurement));
                readings.push(a);

                if second_timer.elapsed().as_secs() >= 1 {
                    let amps = readings.clone();
                    println!("{}", amps.len());

                    let device_token = device_token.clone();

                    readings.clear();
                    thread::spawn(move || {
                        let mut sum = 0f64;
                        for a in &amps {
                            sum += a * a;
                        }
                        let avg = sum / (amps.len() as f64);
                        let sqrt = avg.sqrt();

                        println!("rms: {}a", sqrt);
                        println!("rms: {}w", amps_to_watts(sqrt));

                        let mut buf = Vec::with_capacity(1024);
                        buf.extend(device_token);

                        let mut microamps = U32GroupVarintEncoder::new(& mut buf);
                        for _i in 0..1 {
                            microamps.put_u32((sqrt * 1000000f64) as u32);
                        }
                        microamps.flush();

                        let mut millivolts = U32GroupVarintEncoder::new(& mut buf);
                        for _i in 0..1 {
                            millivolts.put_u32((120 * 1000) as u32);
                        }
                        millivolts.flush();

                        buf.put_u8(1);

                        println!("{}", hex::encode(buf));

                        use embedded_svc::http::{self, client::*, status, Headers, Status};
                        use embedded_svc::io;
                        use esp_idf_svc::http::client::*;

                        let url = String::from("https://api.breker.ninja/readings");

                        info!("About to fetch content from {}", url);

                        let mut client = EspHttpClient::new(&EspHttpClientConfiguration {
                            crt_bundle_attach: Some(esp_idf_sys::esp_crt_bundle_attach),

                            ..Default::default()
                        }).unwrap();

                        let response = client.post(&url).unwrap().submit().unwrap();

                        println!("{}", response.status());

                        // if batch_timer.elapsed().as_secs() >= 10 {

                        //     batch_timer = time::Instant::now();
                        // }
                    });

                    second_timer = time::Instant::now();
                }

                thread::sleep(time::Duration::from_millis(1));
            }
        }
    }
}

fn measure_to_volts(m: i16) -> f64 {
    let max = 1.024f64;
    return max / 2048f64 * (m as f64);
}

fn volts_to_amps(v: f64) -> f64 {
    let max = 30f64;
    return max * v;
}

fn amps_to_watts(a: f64) -> f64 {
    let v = 120f64;
    return v * a;
}

struct U32GroupVarintEncoder<'a> {
    buf: &'a mut Vec<u8>,
    index: usize,
    store: [u32; 4],
    temp: [u8; 17],
}

impl U32GroupVarintEncoder<'_> {
    fn put_u32(&mut self, x: u32) {
        self.store[self.index] = x;
        self.index += 1;
        if self.index == 4 {
            self.flush();
            self.index = 0;
        }
    }

    fn flush(&mut self) {
        // TODO: Is it more efficient to have a tailored version that's called only in Close()?
        // If index is zero, there are no integers to flush
        if self.index == 0 {
            return;
        }
        // In the case we're flushing (the group isn't of size four), the non-values should be zero
        // This ensures the unused entries are all zero in the sizeByte
        for i in self.index..4 {
            self.store[i] = 0u32;
        }
        let mut length = 1;
        // We need to reset the size byte to zero as we only bitwise OR into it, we don't overwrite it
        self.temp[0] = 0;
        for item in self.store.iter().enumerate() {
            let (i, x): (usize, &u32) = item;
            let mut size = 0u8;
            let shifts: [u32; 4] = [24, 16, 8, 0];
            for shift in shifts {
                // Always writes at least one byte -- the first one (shift = 0)
                // Will write more bytes until the rest of the integer is all zeroes
                if (x >> shift) != 0 || shift == 0 {
                    size += 1;
                    self.temp[length] = (x >> shift) as u8;
                    length += 1;
                }
            }
            // We store the size in two of the eight bits in the first byte (sizeByte)
            // 0 means there is one byte in total, hence why we subtract one from size
            self.temp[0] |= (size - 1) << (((3 - i) as u8) * 2);
        }
        // If we're flushing without a full group of four, remove the unused bytes we computed
        // This enables us to realize it's a partial group on decoding thanks to EOF
        if self.index != 4 {
            length -= 4 - self.index;
        }
        self.buf.extend(&self.temp[0..length]);
        return;
    }

    fn new(buf: & mut Vec<u8>) -> U32GroupVarintEncoder {
        return U32GroupVarintEncoder {
            buf,
            index: 0,
            store: [0; 4],
            temp: [0; 17],
        };
    }
}
