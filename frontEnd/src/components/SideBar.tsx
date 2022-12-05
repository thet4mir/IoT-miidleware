import { useState } from 'react'
import { NavLink } from 'react-router-dom'


const SideBar = () => {
  const [open, setOpen] = useState(true)
  const Menus = [
    {title: "Devices", src: "Setting", path: "devices"}
  ]
  return (
    <div className={`${open ? "w-72" : "w-20" } h-screen p-5 pt-8 bg-dark-purple relative duration-300`}>
      <img
        src="/control.png"
        className={`absolute cursor-pointer -right-3 top-9 w-7 rounded-full border-2 border-violet-900 ${!open && "rotate-180"}`}
        onClick={() => {setOpen(!open)}}
      />
      <NavLink to="/">
        <div className="flex gap-x-4 items-center">
          <img src="/logo-white.png" className={`cursor-pointer w-10 duration-500 ${open && "rotate-[360deg]"}`}/>
          <h1
            className={`text-white origin-left font-medium text-xl duration-300 ${!open && "scale-0"}`}
          >
            Andorean
          </h1>
        </div>
      </NavLink>
      <ul className="pt-6">
        {
          Menus.map((menu:any, idx:any) => {
            return (
              <NavLink to={`/${menu.path}`} key={idx}>
                <li
                 key={idx}
                 className="text-gray-300 text-sm flex gap-x-4 items-center cursor-pointer p-2 hover:bg-light-white rounded-md"
                >
                  <img src={`/${menu.src}.png`}/>
                  <span className={`duration-300 ${!open && "scale-0"}`}>{menu.title}</span>
                </li>
              </NavLink>
            )
          })
        }
      </ul>
    </div>
  )
}

export default SideBar
