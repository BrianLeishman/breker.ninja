package main

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Smerity/govarint"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/timestreamwrite"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/xid"
)

func init() {
	r.POST("/readings", hReadingCreate)
	r.GET("/readings/sample", hReadingSample)
}

func hReadingCreate(c *gin.Context) {
	var header struct {
		Time         [4]byte
		UserID       [12]byte
		APIKey       [16]byte
		PlaceID      [12]byte
		BoxID        [12]byte
		DeviceID     [12]byte
		BreakerCount uint8
	}

	b, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrapf(err, "failed to get raw post data").Error()})
		return
	}

	r := bytes.NewReader(b)
	err = binary.Read(r, binary.LittleEndian, &header)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrapf(err, "failed to parse reading header").Error()})
		return
	}

	userID, err := xid.FromBytes(header.UserID[:])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrapf(err, "failed to read userID").Error()})
		return
	}

	placeID, err := xid.FromBytes(header.PlaceID[:])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrapf(err, "failed to read placeID").Error()})
		return
	}

	boxID, err := xid.FromBytes(header.BoxID[:])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrapf(err, "failed to read boxID").Error()})
		return
	}

	deviceID, err := xid.FromBytes(header.DeviceID[:])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrapf(err, "failed to read deviceID").Error()})
		return
	}

	records := make([]*timestreamwrite.Record, 0, header.BreakerCount*2)

	dec := govarint.NewU32GroupVarintDecoder(r)

	for i := 0; i < int(header.BreakerCount); i++ {
		v, err := dec.GetU32()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrapf(err, "failed to read μA number for breaker #%d", i).Error()})
			return
		}

		records = append(records, &timestreamwrite.Record{
			Dimensions: []*timestreamwrite.Dimension{
				{
					Name:  aws.String("breaker"),
					Value: aws.String(deviceID.String() + strconv.Itoa(i)),
				},
			},
			MeasureName:  aws.String("microamps"),
			MeasureValue: aws.String(strconv.Itoa(int(v))),
		})
	}

	for i := 0; i < int(header.BreakerCount); i++ {
		v, err := dec.GetU32()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrapf(err, "failed to read mV number for breaker #%d", i).Error()})
			return
		}

		records = append(records, &timestreamwrite.Record{
			Dimensions: []*timestreamwrite.Dimension{
				{
					Name:  aws.String("breaker"),
					Value: aws.String(deviceID.String() + strconv.Itoa(i)),
				},
			},
			MeasureName:  aws.String("millivolts"),
			MeasureValue: aws.String(strconv.Itoa(int(v))),
		})
	}

	sess, err := session.NewSession()
	writeSvc := timestreamwrite.New(sess)

	currentTimeInSeconds := int64(binary.LittleEndian.Uint32(header.Time[:]))

	writeRecordsInput := &timestreamwrite.WriteRecordsInput{
		DatabaseName: aws.String("breker"),
		TableName:    aws.String("readings"),
		CommonAttributes: &timestreamwrite.Record{
			Dimensions: []*timestreamwrite.Dimension{
				{
					Name:  aws.String("user"),
					Value: aws.String(userID.String()),
				},
				{
					Name:  aws.String("place"),
					Value: aws.String(placeID.String()),
				},
				{
					Name:  aws.String("box"),
					Value: aws.String(boxID.String()),
				},
				{
					Name:  aws.String("device"),
					Value: aws.String(deviceID.String()),
				},
			},
			MeasureValueType: aws.String("BIGINT"),
			Time:             aws.String(strconv.FormatInt(currentTimeInSeconds, 10)),
			TimeUnit:         aws.String("SECONDS"),
		},
		Records: records,
	}

	_, err = writeSvc.WriteRecords(writeRecordsInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrapf(err, "failed to write readings").Error()})
		return
	}
}

func hReadingSample(c *gin.Context) {
	now := time.Now()

	userID := xid.New()

	apiKey, _ := uuid.NewRandom()

	placeID := xid.New()

	boxID := xid.New()

	deviceID := xid.New()

	breakerCount := 48

	var buf bytes.Buffer

	tBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(tBytes, uint32(now.Unix()))

	buf.Write(tBytes)

	buf.Write(userID.Bytes())

	b, _ := apiKey.MarshalBinary()
	buf.Write(b)

	buf.Write(placeID.Bytes())

	buf.Write(boxID.Bytes())

	buf.Write(deviceID.Bytes())

	buf.WriteByte(byte(breakerCount))

	enc := govarint.NewU32GroupVarintEncoder(&buf)

	// μA
	min := 5_000
	max := 100_000_000
	for i := 0; i < breakerCount; i++ {
		rand.Seed(time.Now().UnixNano())
		enc.PutU32(uint32(rand.Intn(max-min) + min))
	}

	// mV
	min = 100_000
	max = 140_000
	for i := 0; i < breakerCount; i++ {
		rand.Seed(time.Now().UnixNano())
		enc.PutU32(uint32(rand.Intn(max-min) + min))
	}

	c.Writer.Write(buf.Bytes())
}
