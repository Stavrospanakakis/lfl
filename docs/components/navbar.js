import { useState } from 'react'

export default function Navbar() {
    const [isClicked, setisClicked] = useState(false)

    const handleClick = () => {
        setisClicked(!isClicked)
    }

    const scrollToInstallation = () => {
      const section = document.getElementById('#installation');
      section?.scrollIntoView( { behavior: 'smooth', block: 'start' } )
      setisClicked(false)
    }

    const scrollToUsage = () => {
      const section = document.getElementById('#usage');
      section?.scrollIntoView( { behavior: 'smooth', block: 'start' } )
      setisClicked(false)
    }

   
    return (
        <header className="fixed bg-black top-0 w-full overflow-hidden sm:flex sm:flex-wrap sm:items-center sm:justify-between px-2 py-3">
        <div className="container mx-auto px-4 sm:flex sm:flex-wrap sm:items-center sm:justify-between">
          <a href="/" className="text-lg font-bold text-white">lfl - Links for Lectures</a>
          <button onClick={handleClick} className="sm:hidden text-white float-right text-xl leading-none px-3 py-1">
             <span className="block relative w-6 rounded-sm bg-white border-white border mt-1"></span>
             <span className="block relative w-6 rounded-sm bg-white mt-1 border"></span>
             <span className="block relative w-6 rounded-sm bg-white mt-1 border"></span>
          </button>
          {isClicked? 
          <div className="block">
              <div onClick={scrollToInstallation} className="text-white pt-2 rounded py-1 block sm:inline-block">Installation</div>
              <div onClick={scrollToUsage} className="text-white rounded py-1 block sm:inline-block">Usage</div>
          </div>
          :
          <div>
              <div onClick={scrollToInstallation} className="text-white px-3 rounded py-1 hidden sm:inline-block cursor-pointer">Installation</div>
              <div onClick={scrollToUsage} className="text-white px-3 rounded py-1 hidden sm:inline-block cursor-pointer">Usage</div>
          </div>
          }
        </div>
      </header>


    )
}