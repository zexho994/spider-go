package entity

import (
	"testing"
)

func TestNode(t *testing.T) {
	type args struct {
		name       string
		attributes *Attributes
	}
	tests := []struct {
		name    string
		args    args
		want    *Node
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "success", args: args{name: "name", attributes: &Attributes{data: map[string]string{}}}, want: nil, wantErr: false},
		{name: "nameBlack", args: args{name: "", attributes: &Attributes{data: map[string]string{}}}, want: nil, wantErr: true},
		{name: "attributeNil", args: args{name: "name", attributes: nil}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewNode(tt.args.name, tt.args.attributes)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_node_updateAttribute(t *testing.T) {
	type fields struct {
		name         string
		attributes   *Attributes
		associated   map[string]*Relation
		beAssociated map[string]*Relation
	}
	type args struct {
		k string
		v string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{name: "success", fields: fields{name: "Node", attributes: NewAttribute("k1", map[string]string{"k1": "v1"})}, args: args{k: "k1", v: "v1"}, want: true, wantErr: false},
		{name: "attrKNil", fields: fields{name: "Node", attributes: NewAttribute("k1", map[string]string{})}, args: args{k: "k1", v: "v1"}, want: false, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Node{
				name:         tt.fields.name,
				attributes:   tt.fields.attributes,
				associated:   tt.fields.associated,
				beAssociated: tt.fields.beAssociated,
			}
			got, err := this.UpdateAttribute(tt.args.k, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("updateAttribute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("updateAttribute() got = %v, want %v", got, tt.want)
			}
		})
	}
}
