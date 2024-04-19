import React from 'react'
import Davis from '@/public/image/Dave.jpg'
import chiks from '@/public/image/chiks.jpg'
import rafa from '@/public/image/rafa.png'



const About = () => {
return (
  <div
    className="flex flex-col items-center justify-center py-40"
    id="aboutus"
  >
    <div className='max-w-[1240px] mx-10 grid md:grid-cols-2 gap-8 gap-y-24'>
      <div data-aos="fade-right">
        <h1 className="text-7xl font-semibold text-transparent bg-clip-text bg-white py-20 text-center md:text-left md:text-9xl">
          ABOUT US.
        </h1>
      </div>

      <div className='w-full bg-[#231B6E]/20 border border-[#2A0E61] flex flex-col p-4 my-4 rounded-lg hover:scale-105 duration-300 items-center' data-aos="zoom-in">
        <img src={Davis.src} alt="" className='rounded-full border-2 border-[#2A0E61] mx-auto mt-[-5rem] object-cover w-[150px] h-[150px]'/> 
        <h2 className='text-2xl w-full font-bold text-center px-5 rounded-full shadow-md bg-[#2A0E61] py-3 my-3'>Davis</h2>
        <p className='text-center pb-10 pt-8 text-white'>
        Hello, I'm Muhammad Davis Adhipramana, you can call me Davis or Dave. I'm an IT student of ITB. currently living in Bandung. I have a wide range of interests, including Web Development, Game Development, Mobile Applications, and Cybersecurity. Currently, the main focus is on mastering Game Development and refining skills in Web Development.
        </p>
      </div>
      <div className='w-full bg-[#231B6E]/20 border border-[#2A0E61] flex flex-col p-4 my-4 rounded-lg hover:scale-105 duration-300 items-center' data-aos="zoom-in">
        <img src={chiks.src} alt="" className='rounded-full border-2 border-[#2A0E61] mx-auto mt-[-5rem] object-cover w-[150px] h-[150px]'/> 
        <h2 className='text-2xl w-full font-bold text-center px-5 rounded-full shadow-md bg-[#2A0E61] py-3 my-3'>Chika</h2>
        <p className='text-center pb-10 pt-8 text-white'>
        Hello, I'm Auralea, you can call me Chika. I'm a sophomore student of Informatics Engineering at Institut Teknologi Bandung. Diving into the world of front-end web development are where I find my passion. I enjoy the creative process of designing user interfaces and bringing them to life through code. Beyond web development, I have a deep curiosity in AI and am eager to delve into its world.
        </p>
      </div>
      <div className='w-full bg-[#231B6E]/20 border border-[#2A0E61] flex flex-col p-4 my-4 rounded-lg hover:scale-105 duration-300 items-center' data-aos="zoom-in">
        <img src={rafa.src} alt="" className='rounded-full border-2 border-[#2A0E61] mx-auto mt-[-5rem] object-cover w-[150px] h-[150px]'/> 
        <h2 className='text-2xl w-full font-bold text-center px-5 rounded-full shadow-md bg-[#2A0E61] py-3 my-3'>Rafa</h2>
        <p className='text-center pb-10 pt-8 text-white'>
        Hello I'm Rafa a sophomore majoring in Informatics Engineering at the Bandung Institute of Technology, I'm a tech enthusiast navigating the ever-evolving landscape of technology. I harbor a deep interest in the gaming industry, constantly exploring its trends and innovations. Currently, the focus is on website development, where I combine a keen eye for design with coding prowess to craft seamless online experiences.
        </p>
      </div>
      
    </div>

  </div>
  )
}

export default About
