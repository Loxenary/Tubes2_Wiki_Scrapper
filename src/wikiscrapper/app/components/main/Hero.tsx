"use client"; // ??

import React from 'react'

const Hero = () => {

  return (
    <div className="flex flex-row items-center justify-center px-20 h-screen w-full z-[20] shadow-inner">
      <div className="h-full w-full flex flex-col gap-5 justify-center  text-center">
        <div className='text-6xl font-semibold text-white w-auto h-auto justify-center items-center' data-aos="fade-down">
          Welcome to <span className='text-transparent bg-clip-text bg-gradient-to-r from-purple-500 to bg-cyan-500'>WikiRace</span>
        </div>

        
        <div data-aos="fade-right">
          <p className='text-lg text-gray-400 my-5'>
            The thrilling online game where you race against the clock to navigate from one Wikipedia page to another using only hyperlinks! Starting from a given Wikipedia page, the challenge is to reach a specific target page by surfing through the interconnected articles. The goal is to find the shortest path possible. Ready to start racing?
          </p>
        </div>

        <div  data-aos="fade-up">
          <a className='py-2 button-primary text-center text-white cursor-pointer rounded-lg'>
            <div className='transform transition duration-300 hover:scale-125 w-[200px] bg-[#221465] rounded-md font-medium my-6 mx-auto py-4 text-white' style={{
              boxShadow: 'inset 0 0 10px rgba(201, 191, 255, 0.5)'
            }}>
              Get Started
            </div>
          </a>
        </div>

      </div>
    </div>
  )
}

export default Hero
