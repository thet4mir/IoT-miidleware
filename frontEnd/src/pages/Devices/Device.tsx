import { useEffect, useState } from 'react'


const Device = (props: any) => {
  const [status, setStatus] = useState(false)
  const { deviceInfo } = props

  return (
    <div className="flex w-1/4 h-[50px] bg-dark-purple p-8 items-center rounded-full">
      <div className="flex w-1/5 items-center">
        <img src="/Setting.png" className="w-[50px]"/>
      </div>
      <div className="flex w-4/3 flex-col gap-y-4 items-center justify-between">
          <div className="flex gap-x-8 items-center justify-between w-full">
            <span className="text-white">
            {deviceInfo[0]}.{deviceInfo[1]}.{deviceInfo[2]}
            </span>
            {
              deviceInfo[3] == "add"
              ?
              <span className={`w-[15px] h-[15px] bg-[#22c55e] rounded-full`}></span>
              :
              <span className={`w-[15px] h-[15px] bg-[#ff0000] rounded-full`}></span>
            }
          </div>
      </div>
    </div>
  )
}

export default Device
