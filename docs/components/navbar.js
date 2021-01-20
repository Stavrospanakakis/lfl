import { useState } from 'react'

export default function Navbar() {
    const [isClicked, setisClicked] = useState(false)

    function handleClick() {
        setisClicked(!isClicked)
    }

    return (
        <header className="fixed bg-black top-0 w-full overflow-hidden sm:flex sm:flex-wrap sm:items-center sm:justify-between px-2 py-3">
        <div className="container mx-auto px-4 sm:flex sm:flex-wrap sm:items-center sm:justify-between">
          <a to="/" className="text-lg font-bold text-white">lfl - Links for Lectures</a>
          <button onClick={handleClick} className="sm:hidden text-white float-right text-xl leading-none px-3 py-1">
             <span className="block relative w-6 rounded-sm bg-white border-white border mt-1"></span>
             <span className="block relative w-6 rounded-sm bg-white mt-1 border"></span>
             <span className="block relative w-6 rounded-sm bg-white mt-1 border"></span>
          </button>
          {isClicked? 
          <div className="block">
              <a href="#installation" className="text-white pt-2 rounded py-1 block sm:inline-block">Installation</a>
              <a href="#usage" className="text-white rounded py-1 block sm:inline-block">Usage</a>
          </div>
          :
          <div>
              <a href="#installation" className="text-white px-3 rounded py-1 hidden sm:inline-block">Installation</a>
              <a href="#usage" className="text-white hover:bg-gray-700 px-3 rounded py-1 hidden sm:inline-block">Usage</a>
          </div>
          }
        </div>
      </header>


    )
}