"use client";
import React, { useState, useRef } from "react";
import ToggleImage from "@/public/image/grad1.png"
import DefaultImage from "@/public/image/grad2.png"
import { StaticImageData } from "next/image";

interface BackgroundOptions {
  default: StaticImageData;
  toggled: StaticImageData;
}

interface ToggleSliderProps {
  height: number;
  width: number;
  defaultValue?: boolean;
  onChange?: (value: boolean) => void;
  background?: BackgroundOptions;
}

const ToggleSlider: React.FC<ToggleSliderProps> = ({
  height,
  width,
  defaultValue = false,
  onChange,
  background,
}) => {
  const [currentImage, setCurrentImage] = useState(background?.default);
  const isTransitioning = useRef(false);
  const [isToggled, setIsToggled] = useState(defaultValue);

  const defaultBackground = {
    default: DefaultImage,
    toggled: ToggleImage,
  };

  const toggle = () => {
    if (isTransitioning.current) return; // Ignore clicks during transition
    const newValue = !isToggled;
    setIsToggled(newValue);
    setCurrentImage(newValue ? background?.toggled : background?.default);
    isTransitioning.current = true; // Set transition state
    if (onChange) {
      onChange(newValue);
    }
  };

  const handleTransitionEnd = () => {
    setCurrentImage(isToggled ? background?.toggled : background?.default);
    isTransitioning.current = false; // Reset transition state
  };

  const translateValue = width - height ; //Transition distance value

  const SetupSlider = (
    <>
      <button
        style={{ height: `${height}px`, width: `${width}px` }}
        className={`border-2 border-[#2A0E61] rounded-full bg-gray-700  relative cursor-pointer overflow-hidden`}
        onClick={toggle}
      >
        {background ? (
          <img src={currentImage?.src} className="w-full h-full" alt="" />
        ) : null}
        <div
          onTransitionEnd={handleTransitionEnd}
          style={{
            width: `${height}px`,
            transform: `${isToggled ? `translateX(${translateValue}px)` : ``}`,
            transition: ".4s",
          }}
          className={`h-full rounded-full bg-white shadow-xl transform  transition-transform absolute top-0 left-0 right-0 bottom-0 ${
            isToggled ? defaultBackground.toggled : defaultBackground.default
          }`}
        ></div>
      </button>
    </>
  );

  return SetupSlider;
};

export default ToggleSlider;
