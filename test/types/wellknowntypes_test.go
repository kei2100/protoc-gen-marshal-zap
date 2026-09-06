package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	wktTime     = time.Date(2024, 1, 15, 10, 30, 0, 123456789, time.UTC)
	wktDuration = 90*time.Second + 500*time.Millisecond
)

func newWellKnownTypes(t *testing.T) *WellKnownTypes {
	t.Helper()
	anyVal, err := anypb.New(&OtherType3{Val: "any"})
	require.NoError(t, err)
	return &WellKnownTypes{
		TimestampVal:   timestamppb.New(wktTime),
		DurationVal:    durationpb.New(wktDuration),
		BoolValueVal:   wrapperspb.Bool(true),
		StringValueVal: wrapperspb.String("string"),
		BytesValueVal:  wrapperspb.Bytes([]byte{1, 2, 3}),
		Int32ValueVal:  wrapperspb.Int32(-1),
		Int64ValueVal:  wrapperspb.Int64(-2),
		Uint32ValueVal: wrapperspb.UInt32(1),
		Uint64ValueVal: wrapperspb.UInt64(2),
		FloatValueVal:  wrapperspb.Float(0.5),
		DoubleValueVal: wrapperspb.Double(0.25),
		FieldMaskVal:   &fieldmaskpb.FieldMask{Paths: []string{"foo", "bar.baz"}},
		EmptyVal:       &emptypb.Empty{},

		AnyVal:       anyVal,
		StructVal:    &structpb.Struct{Fields: map[string]*structpb.Value{"foo": structpb.NewStringValue("bar")}},
		ValueVal:     structpb.NewNumberValue(1),
		ListValueVal: &structpb.ListValue{Values: []*structpb.Value{structpb.NewBoolValue(true)}},

		NotPresentTimestampVal:   nil,
		NotPresentStringValueVal: nil,
		NotPresentEmptyVal:       nil,

		MaskedTimestampVal: timestamppb.New(wktTime),

		// Repeated (including nil elements)
		RepeatedTimestampVal:   []*timestamppb.Timestamp{timestamppb.New(wktTime), nil},
		RepeatedDurationVal:    []*durationpb.Duration{durationpb.New(wktDuration), nil},
		RepeatedBoolValueVal:   []*wrapperspb.BoolValue{wrapperspb.Bool(true), nil},
		RepeatedStringValueVal: []*wrapperspb.StringValue{wrapperspb.String("string"), nil},
		RepeatedBytesValueVal:  []*wrapperspb.BytesValue{wrapperspb.Bytes([]byte{1, 2, 3}), nil},
		RepeatedInt32ValueVal:  []*wrapperspb.Int32Value{wrapperspb.Int32(-1), nil},
		RepeatedInt64ValueVal:  []*wrapperspb.Int64Value{wrapperspb.Int64(-2), nil},
		RepeatedUint32ValueVal: []*wrapperspb.UInt32Value{wrapperspb.UInt32(1), nil},
		RepeatedUint64ValueVal: []*wrapperspb.UInt64Value{wrapperspb.UInt64(2), nil},
		RepeatedFloatValueVal:  []*wrapperspb.FloatValue{wrapperspb.Float(0.5), nil},
		RepeatedDoubleValueVal: []*wrapperspb.DoubleValue{wrapperspb.Double(0.25), nil},
		RepeatedFieldMaskVal:   []*fieldmaskpb.FieldMask{{Paths: []string{"foo", "bar.baz"}}, nil},
		RepeatedEmptyVal:       []*emptypb.Empty{{}, nil},

		// Map (single entry each so that JSON output is deterministic)
		MapTimestampVal:   map[string]*timestamppb.Timestamp{"k": timestamppb.New(wktTime)},
		MapDurationVal:    map[string]*durationpb.Duration{"k": durationpb.New(wktDuration)},
		MapBoolValueVal:   map[string]*wrapperspb.BoolValue{"k": wrapperspb.Bool(true)},
		MapStringValueVal: map[string]*wrapperspb.StringValue{"k": wrapperspb.String("string")},
		MapBytesValueVal:  map[string]*wrapperspb.BytesValue{"k": wrapperspb.Bytes([]byte{1, 2, 3})},
		MapInt32ValueVal:  map[string]*wrapperspb.Int32Value{"k": wrapperspb.Int32(-1)},
		MapInt64ValueVal:  map[string]*wrapperspb.Int64Value{"k": wrapperspb.Int64(-2)},
		MapUint32ValueVal: map[string]*wrapperspb.UInt32Value{"k": wrapperspb.UInt32(1)},
		MapUint64ValueVal: map[string]*wrapperspb.UInt64Value{"k": wrapperspb.UInt64(2)},
		MapFloatValueVal:  map[string]*wrapperspb.FloatValue{"k": wrapperspb.Float(0.5)},
		MapDoubleValueVal: map[string]*wrapperspb.DoubleValue{"k": wrapperspb.Double(0.25)},
		MapFieldMaskVal:   map[string]*fieldmaskpb.FieldMask{"k": {Paths: []string{"foo", "bar.baz"}}},
		MapEmptyVal:       map[string]*emptypb.Empty{"nil": nil},

		OneofVal: &WellKnownTypes_OneofTimestampVal{OneofTimestampVal: timestamppb.New(wktTime)},

		OptionalTimestampVal:           timestamppb.New(wktTime),
		OptionalNotPresentTimestampVal: nil,
	}
}

func TestWellKnownTypes_MarshalLogObject(t *testing.T) {
	m := newWellKnownTypes(t)
	enc := zapcore.NewMapObjectEncoder()
	require.NoError(t, m.MarshalLogObject(enc))

	assert.EqualValues(t, map[string]interface{}{
		"timestamp_val":    wktTime,
		"duration_val":     wktDuration,
		"bool_value_val":   true,
		"string_value_val": "string",
		"bytes_value_val":  []byte{1, 2, 3},
		"int32_value_val":  int32(-1),
		"int64_value_val":  int64(-2),
		"uint32_value_val": uint32(1),
		"uint64_value_val": uint64(2),
		"float_value_val":  float32(0.5),
		"double_value_val": float64(0.25),
		"field_mask_val":   []interface{}{"foo", "bar.baz"},
		"empty_val":        map[string]interface{}{},

		// generic message fallback: AddReflected keeps the proto message as is
		"any_val":        m.AnyVal,
		"struct_val":     m.StructVal,
		"value_val":      m.ValueVal,
		"list_value_val": m.ListValueVal,

		"masked_timestamp_val": "[MASKED]",

		"repeated_timestamp_val":    []interface{}{wktTime, nil},
		"repeated_duration_val":     []interface{}{wktDuration, nil},
		"repeated_bool_value_val":   []interface{}{true, nil},
		"repeated_string_value_val": []interface{}{"string", nil},
		"repeated_bytes_value_val":  []interface{}{"\x01\x02\x03", nil}, // MapObjectEncoder.AppendByteString stores a string
		"repeated_int32_value_val":  []interface{}{int32(-1), nil},
		"repeated_int64_value_val":  []interface{}{int64(-2), nil},
		"repeated_uint32_value_val": []interface{}{uint32(1), nil},
		"repeated_uint64_value_val": []interface{}{uint64(2), nil},
		"repeated_float_value_val":  []interface{}{float32(0.5), nil},
		"repeated_double_value_val": []interface{}{float64(0.25), nil},
		"repeated_field_mask_val":   []interface{}{[]interface{}{"foo", "bar.baz"}, nil},
		"repeated_empty_val":        []interface{}{map[string]interface{}{}, nil},

		"map_timestamp_val":    map[string]interface{}{"k": wktTime},
		"map_duration_val":     map[string]interface{}{"k": wktDuration},
		"map_bool_value_val":   map[string]interface{}{"k": true},
		"map_string_value_val": map[string]interface{}{"k": "string"},
		"map_bytes_value_val":  map[string]interface{}{"k": []byte{1, 2, 3}},
		"map_int32_value_val":  map[string]interface{}{"k": int32(-1)},
		"map_int64_value_val":  map[string]interface{}{"k": int64(-2)},
		"map_uint32_value_val": map[string]interface{}{"k": uint32(1)},
		"map_uint64_value_val": map[string]interface{}{"k": uint64(2)},
		"map_float_value_val":  map[string]interface{}{"k": float32(0.5)},
		"map_double_value_val": map[string]interface{}{"k": float64(0.25)},
		"map_field_mask_val":   map[string]interface{}{"k": []interface{}{"foo", "bar.baz"}},
		"map_empty_val":        map[string]interface{}{"nil": nil},

		"oneof_timestamp_val": wktTime,

		"optional_timestamp_val": wktTime,
	}, enc.Fields)

	// Explicit presence: not present fields must be omitted.
	assert.NotContains(t, enc.Fields, "not_present_timestamp_val")
	assert.NotContains(t, enc.Fields, "not_present_string_value_val")
	assert.NotContains(t, enc.Fields, "not_present_empty_val")
	assert.NotContains(t, enc.Fields, "oneof_duration_val")
	assert.NotContains(t, enc.Fields, "optional_not_present_timestamp_val")
}

// TestWellKnownTypes_JSONEncoder verifies the actual output of zapcore's JSON encoder,
// which is what users see in their logs.
func TestWellKnownTypes_JSONEncoder(t *testing.T) {
	m := &WellKnownTypes{
		TimestampVal:         timestamppb.New(wktTime),
		DurationVal:          durationpb.New(wktDuration),
		StringValueVal:       wrapperspb.String("string"),
		BytesValueVal:        wrapperspb.Bytes([]byte{1, 2, 3}),
		FieldMaskVal:         &fieldmaskpb.FieldMask{Paths: []string{"foo", "bar.baz"}},
		EmptyVal:             &emptypb.Empty{},
		StructVal:            &structpb.Struct{Fields: map[string]*structpb.Value{"foo": structpb.NewStringValue("bar")}},
		RepeatedTimestampVal: []*timestamppb.Timestamp{timestamppb.New(wktTime), nil},
		RepeatedEmptyVal:     []*emptypb.Empty{{}, nil},
		MapTimestampVal:      map[string]*timestamppb.Timestamp{"k": timestamppb.New(wktTime)},
		MapEmptyVal:          map[string]*emptypb.Empty{"nil": nil},
	}

	encode := func(t *testing.T, cfg zapcore.EncoderConfig) string {
		t.Helper()
		// suppress entry-level keys
		cfg.TimeKey = ""
		cfg.LevelKey = ""
		cfg.MessageKey = ""
		buf, err := zapcore.NewJSONEncoder(cfg).EncodeEntry(zapcore.Entry{}, []zapcore.Field{zap.Object("m", m)})
		require.NoError(t, err)
		return buf.String()
	}

	t.Run("ISO8601TimeEncoder+StringDurationEncoder", func(t *testing.T) {
		cfg := zap.NewDevelopmentEncoderConfig()
		cfg.EncodeTime = zapcore.ISO8601TimeEncoder
		cfg.EncodeDuration = zapcore.StringDurationEncoder
		assert.Equal(t, `{"m":{`+
			`"timestamp_val":"2024-01-15T10:30:00.123Z",`+
			`"duration_val":"1m30.5s",`+
			`"string_value_val":"string",`+
			`"bytes_value_val":"AQID",`+
			`"field_mask_val":["foo","bar.baz"],`+
			`"empty_val":{},`+
			`"struct_val":{"foo":"bar"},`+
			`"masked_timestamp_val":"[MASKED]",`+
			`"repeated_timestamp_val":["2024-01-15T10:30:00.123Z",null],`+
			`"repeated_duration_val":[],`+
			`"repeated_bool_value_val":[],`+
			`"repeated_string_value_val":[],`+
			`"repeated_bytes_value_val":[],`+
			`"repeated_int32_value_val":[],`+
			`"repeated_int64_value_val":[],`+
			`"repeated_uint32_value_val":[],`+
			`"repeated_uint64_value_val":[],`+
			`"repeated_float_value_val":[],`+
			`"repeated_double_value_val":[],`+
			`"repeated_field_mask_val":[],`+
			`"repeated_empty_val":[{},null],`+
			`"map_timestamp_val":{"k":"2024-01-15T10:30:00.123Z"},`+
			`"map_duration_val":{},`+
			`"map_bool_value_val":{},`+
			`"map_string_value_val":{},`+
			`"map_bytes_value_val":{},`+
			`"map_int32_value_val":{},`+
			`"map_int64_value_val":{},`+
			`"map_uint32_value_val":{},`+
			`"map_uint64_value_val":{},`+
			`"map_float_value_val":{},`+
			`"map_double_value_val":{},`+
			`"map_field_mask_val":{},`+
			`"map_empty_val":{"nil":null}`+
			`}}`+"\n", encode(t, cfg))
	})

	t.Run("EpochTimeEncoder+SecondsDurationEncoder", func(t *testing.T) {
		// zap.NewProductionEncoderConfig() defaults
		out := encode(t, zap.NewProductionEncoderConfig())
		assert.Contains(t, out, `"timestamp_val":1705314600.1234567`)
		assert.Contains(t, out, `"duration_val":90.5,`)
		assert.Contains(t, out, `"map_timestamp_val":{"k":1705314600.1234567`)
	})
}
