import { createContext, useContext, useState } from "react";
import { IOutputContext } from "@/app/Output/outputData";
import PathInterface from "@/app/Output/PathData";

///<summary>
/// This Context is used to hold the information for the output data from the response of the server
///</summary>


const OutputContext = createContext<IOutputContext | undefined>(undefined);

export const useOutputContext = () => {
  const context = useContext(OutputContext);
  if (!context) {
    throw new Error(
      "useOutputContext must be used within a OutputContextProvider"
    );
  }
  return context;
};

interface OutputContextProvider {
  children: React.ReactNode;
}

export const OutputContextProvider: React.FC<OutputContextProvider> = ({
  children,
}) => {
  const [checkcount, setCheckcount] = useState("");
  const [numpassed, setNumpassed] = useState("");
  const [time, setTime] = useState("");
  const [listPath, setListPath] = useState<PathInterface[]>([]);

  function setOutputData (checkcount: string, numpassed : string, time : string, pathList : PathInterface[]){
    setCheckcount(checkcount);
    setNumpassed(numpassed);
    setTime(time);
    setListPath(pathList);
  }

  return(
    <OutputContext.Provider value={{ checkcount, numpassed, time, listPath, setOutputData }}>
      {children}
    </OutputContext.Provider>
  )
};
