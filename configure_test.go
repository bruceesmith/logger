package logger

import (
	"bytes"
	"io"
	"testing"
)

func TestConfigure(t *testing.T) {
	type args struct {
		setting []ConfigSetting
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "no-such-log",
			args: args{
				setting: []ConfigSetting{
					{
						AppliesTo: 77,
						Key:       99,
						Value:     "any",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "no-such-key",
			args: args{
				setting: []ConfigSetting{
					{
						AppliesTo: Norm,
						Key:       99,
						Value:     "any",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "bad-destination",
			args: args{
				setting: []ConfigSetting{
					{
						AppliesTo: Norm,
						Key:       DestinationSetting,
						Value:     "any",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "bad-format",
			args: args{
				setting: []ConfigSetting{
					{
						AppliesTo: Norm,
						Key:       FormatSetting,
						Value:     "any",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "bad-omittime",
			args: args{
				setting: []ConfigSetting{
					{
						AppliesTo: Norm,
						Key:       OmitTimeSetting,
						Value:     "any",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "bad-omittime",
			args: args{
				setting: []ConfigSetting{
					{
						AppliesTo: Norm,
						Key:       OmitTimeSetting,
						Value:     "any",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "destination",
			args: args{
				setting: []ConfigSetting{
					{
						AppliesTo: Norm,
						Key:       DestinationSetting,
						Value:     &bytes.Buffer{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "format",
			args: args{
				setting: []ConfigSetting{
					{
						AppliesTo: Norm,
						Key:       FormatSetting,
						Value:     Text,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "omittime",
			args: args{
				setting: []ConfigSetting{
					{
						AppliesTo: Norm,
						Key:       OmitTimeSetting,
						Value:     true,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Configure(tt.args.setting...); (err != nil) != tt.wantErr {
				t.Errorf("Configure() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_formats(t *testing.T) {
	type args struct {
		log LogID
		f   Format
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "norm",
			args: args{
				log: Norm,
				f:   JSON,
			},
		},
		{
			name: "tracy",
			args: args{
				log: Tracy,
				f:   JSON,
			},
		},
	}
	for _, tt := range tests {
		save := config
		t.Run(tt.name, func(t *testing.T) {
			formats(tt.args.log, tt.args.f)
			if tt.args.log == Norm && config.Normal.Format != tt.args.f {
				t.Errorf("formats() got = %v want =%v", config.Normal.Format, tt.args.f)
			}
			if tt.args.log == Tracy && config.Trace.Format != tt.args.f {
				t.Errorf("formats() got = %v want =%v", config.Trace.Format, tt.args.f)
			}
		})
		config = save
	}
}

func Test_normalFormat(t *testing.T) {
	type args struct {
		f Format
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "json",
			args: args{
				f: JSON,
			},
		},
		{
			name: "text",
			args: args{
				f: Text,
			},
		},
	}
	for _, tt := range tests {
		save := config
		t.Run(tt.name, func(t *testing.T) {
			normalFormat(tt.args.f)
			if config.Normal.Format != tt.args.f {
				t.Errorf("normalFormat() got = %v want =%v", config.Normal.Format, tt.args.f)
			}
		})
		config = save
	}
}

func Test_traceFormat(t *testing.T) {
	type args struct {
		f Format
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "json",
			args: args{
				f: JSON,
			},
		},
		{
			name: "text",
			args: args{
				f: Text,
			},
		},
	}
	for _, tt := range tests {
		save := config
		t.Run(tt.name, func(t *testing.T) {
			traceFormat(tt.args.f)
			if config.Trace.Format != tt.args.f {
				t.Errorf("traceFormat() got = %v want =%v", config.Trace.Format, tt.args.f)
			}
		})
		config = save
	}
}

func Test_destination(t *testing.T) {
	type args struct {
		log LogID
		w   io.Writer
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "norm",
			args: args{
				log: Norm,
				w:   &bytes.Buffer{},
			},
		},
		{
			name: "tracy",
			args: args{
				log: Tracy,
				w:   &bytes.Buffer{},
			},
		},
	}
	for _, tt := range tests {
		save := config
		t.Run(tt.name, func(t *testing.T) {
			destination(tt.args.log, tt.args.w)
			if tt.args.log == Norm && config.Normal.Destination != tt.args.w {
				t.Errorf("destination() got = %v want =%v", config.Normal.Destination, tt.args.w)
			}
			if tt.args.log == Tracy && config.Trace.Destination != tt.args.w {
				t.Errorf("destination() got = %v want =%v", config.Trace.Destination, tt.args.w)
			}
		})
		config = save
	}
}

func Test_normalDestination(t *testing.T) {
	type args struct {
		f Format
		w io.Writer
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "json",
			args: args{
				f: JSON,
				w: &bytes.Buffer{},
			},
		},
		{
			name: "text",
			args: args{
				f: Text,
				w: &bytes.Buffer{},
			},
		},
	}
	for _, tt := range tests {
		save := config
		config.Normal.Format = tt.args.f
		t.Run(tt.name, func(t *testing.T) {
			normalDestination(tt.args.w)
			if config.Normal.Destination != tt.args.w {
				t.Errorf("normalDestination() got = %v want =%v", config.Normal.Destination, tt.args.w)
			}
		})
		config = save
	}
}

func Test_traceDestination(t *testing.T) {
	type args struct {
		f Format
		w io.Writer
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "json",
			args: args{
				f: JSON,
				w: &bytes.Buffer{},
			},
		},
		{
			name: "text",
			args: args{
				f: Text,
				w: &bytes.Buffer{},
			},
		},
	}
	for _, tt := range tests {
		save := config
		config.Trace.Format = tt.args.f
		t.Run(tt.name, func(t *testing.T) {
			traceDestination(tt.args.w)
			if config.Trace.Destination != tt.args.w {
				t.Errorf("traceDestination() got = %v want =%v", config.Trace.Destination, tt.args.w)
			}
		})
		config = save
	}
}

func Test_omitTime(t *testing.T) {
	type args struct {
		log  LogID
		omit bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "norm",
			args: args{
				log:  Norm,
				omit: true,
			},
		},
		{
			name: "tracy",
			args: args{
				log:  Tracy,
				omit: true,
			},
		},
	}
	for _, tt := range tests {
		save := config
		t.Run(tt.name, func(t *testing.T) {
			omitTime(tt.args.log, tt.args.omit)
			if tt.args.log == Norm && config.Normal.OmitTime != tt.args.omit {
				t.Errorf("omitTime() got = %v want =%v", config.Normal.OmitTime, tt.args.omit)
			}
			if tt.args.log == Tracy && config.Trace.OmitTime != tt.args.omit {
				t.Errorf("omitTime() got = %v want =%v", config.Trace.OmitTime, tt.args.omit)
			}
		})
		config = save
	}
}
