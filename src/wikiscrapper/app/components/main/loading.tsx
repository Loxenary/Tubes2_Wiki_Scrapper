"use client"
import { useEffect, useState } from "react";

interface ISetupLoading {
  isLoading: boolean;
}

export const LoadingBar: React.FC<ISetupLoading> = ({ isLoading }) => {
  const [point, setPoint] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setPoint((prevPoint) => (prevPoint + 1) % 3);
    }, 500)
    return () => {
      clearInterval(interval)
      console.log("Loading");
    };
  }, [isLoading]);

  return <ShowLoadingBar isLoading={isLoading} point={point} />;
};

export const ShowLoadingBar = ({
  isLoading,
  point,
}: {
  isLoading: boolean;
  point: number;
}) => {
  return (
    <div className="text-lg justify-center items-center">
      {isLoading ? `Loading${".".repeat(point)}` : null}
    </div>
  );
};
