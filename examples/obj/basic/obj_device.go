package basic

import (
	. "github.com/zubairhamed/go-lwm2m/api"
	"github.com/zubairhamed/go-lwm2m/core"
	"github.com/zubairhamed/go-lwm2m/core/response"
	"github.com/zubairhamed/go-lwm2m/core/values"
	"github.com/zubairhamed/go-lwm2m/objects/oma"
	"time"
)

type Device struct {
	Model 			ObjectModel
	currentTime		time.Time
	utcOffset		string
	timeZone		string
}

func (o *Device) OnExecute(instanceId int, resourceId int, req Request) Response {
	return response.Changed()
}

func (o *Device) OnCreate(instanceId int, resourceId int, req Request) Response {
	return response.Created()
}

func (o *Device) OnDelete(instanceId int, req Request) Response {
	return response.Deleted()
}

func (o *Device) OnRead(instanceId int, resourceId int, req Request) Response {
	if resourceId == -1 {
		// Read Object Instance
	} else {
		// Read Resource Instance
		var val ResponseValue

		resource := o.Model.GetResource(resourceId)
		switch resourceId {
		case 0:
			val = values.String("Open Mobile Alliance")
			break

		case 1:
			val = values.String("Lightweight M2M Client")
			break

		case 2:
			val = values.String("345000123")
			break

		case 3:
			val = values.String("1.0")
			break

		case 6:
			val, _ = core.TlvPayloadFromIntResource(resource, []int{1, 5})
			break

		case 7:
			val, _ = core.TlvPayloadFromIntResource(resource, []int{3800, 5000})
			break

		case 8:
			val, _ = core.TlvPayloadFromIntResource(resource, []int{125, 900})
			break

		case 9:
			val = values.Integer(100)
			break

		case 10:
			val = values.Integer(15)
			break

		case 11:
			val, _ = core.TlvPayloadFromIntResource(resource, []int{0})
			break

		case 13:
			val = values.Time(o.currentTime)
			break

		case 14:
			val = values.String(o.utcOffset)
			break

		case 15:
			val = values.String(o.timeZone)
			break

		case 16:
			val = values.String("U")
			break

		default:
			break
		}
		return response.Content(val)
	}
	return response.NotFound()
}

func (o *Device) OnWrite(instanceId int, resourceId int, req Request) Response {
	val := req.GetMessage().Payload

	switch resourceId {
	case 13:
		break

	case 14:
		o.utcOffset = val.String()
		break

	case 15:
		o.timeZone = val.String()
		break

	default:
		return response.NotFound()
	}
	return response.Changed()
}

func (o *Device) Reboot() ResponseValue {
	return values.Empty()
}

func (o *Device) FactoryReset() ResponseValue {
	return values.Empty()
}

func (o *Device) ResetErrorCode() string {
	return ""
}

func NewExampleDeviceObject(reg Registry) *Device {
	return &Device{
		Model: reg.GetModel(oma.OBJECT_LWM2M_DEVICE),
		currentTime: time.Unix(1367491215, 0),
		utcOffset: "+02:00",
		timeZone: "+02:00",
	}
}
