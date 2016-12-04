// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package member

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

// Attributes:
//  - _id
//  - LineID
//  - Name
//  - Pic
//  - Message
//  - Day
type Member struct {
	_id     string `thrift:"_id,1" json:"_id"`
	LineID  string `thrift:"line_id,2" json:"line_id"`
	Name    string `thrift:"name,3" json:"name"`
	Pic     string `thrift:"pic,4" json:"pic"`
	Message string `thrift:"message,5" json:"message"`
	Day     string `thrift:"day,6" json:"day"`
}

func NewMember() *Member {
	return &Member{}
}

func (p *Member) Get_id() string {
	return p._id
}

func (p *Member) GetLineID() string {
	return p.LineID
}

func (p *Member) GetName() string {
	return p.Name
}

func (p *Member) GetPic() string {
	return p.Pic
}

func (p *Member) GetMessage() string {
	return p.Message
}

func (p *Member) GetDay() string {
	return p.Day
}
func (p *Member) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Member) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p._id = v
	}
	return nil
}

func (p *Member) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.LineID = v
	}
	return nil
}

func (p *Member) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Name = v
	}
	return nil
}

func (p *Member) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Pic = v
	}
	return nil
}

func (p *Member) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Message = v
	}
	return nil
}

func (p *Member) readField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.Day = v
	}
	return nil
}

func (p *Member) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Member"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Member) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("_id", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:_id: ", p), err)
	}
	if err := oprot.WriteString(string(p._id)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T._id (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:_id: ", p), err)
	}
	return err
}

func (p *Member) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("line_id", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:line_id: ", p), err)
	}
	if err := oprot.WriteString(string(p.LineID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.line_id (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:line_id: ", p), err)
	}
	return err
}

func (p *Member) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("name", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:name: ", p), err)
	}
	if err := oprot.WriteString(string(p.Name)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.name (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:name: ", p), err)
	}
	return err
}

func (p *Member) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("pic", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:pic: ", p), err)
	}
	if err := oprot.WriteString(string(p.Pic)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.pic (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:pic: ", p), err)
	}
	return err
}

func (p *Member) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("message", thrift.STRING, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:message: ", p), err)
	}
	if err := oprot.WriteString(string(p.Message)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.message (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:message: ", p), err)
	}
	return err
}

func (p *Member) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("day", thrift.STRING, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:day: ", p), err)
	}
	if err := oprot.WriteString(string(p.Day)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.day (6) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:day: ", p), err)
	}
	return err
}

func (p *Member) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Member(%+v)", *p)
}

// Attributes:
//  - Day
//  - Color
//  - Members
type ResultDay struct {
	Day     string    `thrift:"day,1" json:"day"`
	Color   string    `thrift:"color,2" json:"color"`
	Members []*Member `thrift:"members,3" json:"members"`
}

func NewResultDay() *ResultDay {
	return &ResultDay{}
}

func (p *ResultDay) GetDay() string {
	return p.Day
}

func (p *ResultDay) GetColor() string {
	return p.Color
}

func (p *ResultDay) GetMembers() []*Member {
	return p.Members
}
func (p *ResultDay) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ResultDay) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Day = v
	}
	return nil
}

func (p *ResultDay) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Color = v
	}
	return nil
}

func (p *ResultDay) readField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*Member, 0, size)
	p.Members = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &Member{}
		if err := _elem0.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
		}
		p.Members = append(p.Members, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *ResultDay) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ResultDay"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ResultDay) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("day", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:day: ", p), err)
	}
	if err := oprot.WriteString(string(p.Day)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.day (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:day: ", p), err)
	}
	return err
}

func (p *ResultDay) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("color", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:color: ", p), err)
	}
	if err := oprot.WriteString(string(p.Color)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.color (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:color: ", p), err)
	}
	return err
}

func (p *ResultDay) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("members", thrift.LIST, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:members: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Members)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Members {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:members: ", p), err)
	}
	return err
}

func (p *ResultDay) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ResultDay(%+v)", *p)
}
