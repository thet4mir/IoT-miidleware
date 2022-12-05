import { useEffect, useState } from 'react'
import Device from './Device'
import { DeviceServiceClient } from './devicepb/device_grpc_web_pb'
const {DeviceListRequest, DeviceListResponse} = require('./devicepb/device_pb.js')


var client = new DeviceServiceClient('http://127.0.0.1:8000', null, null)

const DeviceList = () => {

  const [devices, setDevices] = useState([])

  const getDevices = () => {
    var request = new DeviceListRequest()
    var stream = client.deviceList(request, {})
    stream.on('data', (response: any) => {
      setDevices(response.getDeviceList())
    })
  }
  useEffect(() => {
    getDevices()
  }, [])

  return (
    <div className="flex p-8 w-full gap-x-4">
      {
        devices &&
          devices.map((device: any) => {
            return <Device deviceInfo={device.array} key={device}/>
          })
      }
    </div>
  )
}

export default DeviceList
