import React from 'react'
import logows from '@/public/image/logowiki.png'


const Navbar = () => {
  return (
    <div className='w-full h-[65px] fixed top-0 shadow-lg shadow-[#2A0E61]/50 bg-[#03001417] backdrop-blur-md z-50 px-5'>
        <div className='w-full h-full flex flex-row items-center justify-between m-auto px-[10px]'>
            <a href="/" className='h-auto w-auto flex flex-row items-center'>
                <img src={logows.src} alt='logo' width={55} height={60} className="cursor-pointer h-[50px] hover:animate-slowspin" />
            </a>
            <div className='w-[300px] h-full flex flex-row items-center'>
                <div className='flex items-center justify-between w-full h-auto border border-[#7042f861] bg-[#0300145e] px-[20px] py-[10px] rounded-full text-gray-200'>
                    <a href="/" className='cursor-pointer hover:text-purple-800'>Home</a>
                    <a href="#information" className='cursor-pointer hover:text-purple-800'>Information</a>
                    <a href="#aboutus" className='cursor-pointer hover:text-purple-800'>About Us</a>
                </div>

            </div>
        </div>





    </div>
  )
}

export default Navbar
