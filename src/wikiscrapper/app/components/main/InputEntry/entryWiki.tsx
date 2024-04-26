"use client";
import { showToast } from "@/app/components/main/toast";
import { WikipediaExistChecker } from "@/pages/api/wiki";
import {
  SearchWikiInterface,
  useWikiSearchContext,
} from "@/Context/SearchContext";
import { useContext, useState, ChangeEvent, FormEvent } from "react";
import { BoolOutputSetup, ISetupOutputPage } from "@/app/page";
import { IOutputContext } from "../Output/outputData";
import { useOutputContext } from "@/Context/OutputContext";
import { LoadingBar } from "@/app/components/main/loading";
import Autocomplete from "./autocomplete";
const EntryWiki = () => {
  // Save Data for Input Textarea
  const [formValue, setFormValue] = useState({
    FROM: "",
    TO: "",
  });

  const { setOutputState } = useContext<ISetupOutputPage>(BoolOutputSetup);

  // Save Data for loading animation logic
  const [isLoading, setisLoading] = useState(false);

  // Context that saves the output data, used in output page later
  const { setOutputData }: IOutputContext = useOutputContext();

  // Save the data of the algorithm that is used
  const { Algorithm }: SearchWikiInterface = useWikiSearchContext();

  // Used to control the state of the autocomplete when user click on the input field
  const [isFromAutocompleteOpen, setIsFromAutocompleteOpen] = useState(false);
  const [isToAutocompleteOpen, setIsToAutocompleteOpen] = useState(false);

  // used to save the state of current input field
  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { id, value } = event.target;
    //TODO: handle autocomplete from wikipedia api
    setFormValue((prevState) => ({
      ...prevState,
      [id]: value,
    }));
  };

  const handleGetApi = async() =>{
    try{
      const url = "/api/getData";
      const res = await fetch(url, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        }
      })
      if(res.status===204){
        return "null"
      }
      else if(res.status === 200){
        const output = await res.json();
        return output
      }
      if(!res.ok){
        throw new Error("Error Fetching");
      }
    }catch(err){
      showToast("Error Fetching","error")
      setisLoading(false)
      setOutputState(false)
      throw err
    }
  }

  const handleBackendPolling = async () => {
    try{
      const data = await handleGetApi();
      if(data != "null"){
        showToast(JSON.stringify(data),"info");
        handleOutputData(data);
        setOutputState(true);
        setisLoading(false);
        return;
      }
      setTimeout(handleBackendPolling,3000);
    }catch(e){
      throw e
    }

  }

  // used to save the data of the response from backend into an output context
  function handleOutputData(data: any) {
    const { checkcount, listPath, numpassed, time } = data;
    setOutputData(checkcount, numpassed, time, listPath);
  }

  // insert data of algorithm into data links
  function handleInputData(linksdata: any) {
    const dataUsed = {
      FROM: linksdata.FROM,
      TO: linksdata.TO,
      algorithm: Algorithm,
    };
    return dataUsed;
  }

  // handle any input field edge case e.g: empty input or both field has the same input
  function handleInputEdgeCases() {
    if (formValue.FROM === formValue.TO) {
      throw new Error("The Data of From and To are the Same");
    } else if (formValue.FROM.length === 0 || formValue.TO.length === 0) {
      throw new Error("Please fill all the input fields");
    }
  }

  // handle api post to backend return a json object containing output data
  const HandlePostAPI = async (data: any) => {
    const res = await fetch("/api/postData", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });

    if (!res.ok) {
      throw new Error("failed to fetch");
    }
    showToast("Data is Successfully updated","success");
  };

  // formatting data used later in creating wikipedia links
  function removeSpace(data: string) {
    return data.replace(/\s+/g, "_");
  }

  // convert a title into a wikipedia link
  function convertLink(data: string) {
    return `https://en.wikipedia.org/wiki/${removeSpace(data)}`;
  }

  function passedDataLinkConverter(data: string){
    return `/wiki/${removeSpace(data)}`;
  }

  // handle any invalid links, if valid, convert the data to a link and save it into links state value
  const handleTitleToLinks = async () => {
    const { FROM, TO } = formValue;
    const formattedFrom = removeSpace(FROM);
    const formattedTo = removeSpace(TO);

    try {
      const responseFrom = await WikipediaExistChecker(FROM);


      if (!responseFrom.includes(FROM)) {
        throw new Error(`${FROM} title can not found on Wikipedia`);
      }

      const responseTo = await WikipediaExistChecker(TO);
      if (!responseTo.includes(TO)) {
        throw new Error(`${TO} title can not found on Wikipedia`);
      }
      const data = {
        FROM: passedDataLinkConverter(formattedFrom),
        TO: passedDataLinkConverter(formattedTo),
      };
      const out = handleInputData(data);
      return out;
    } catch (error) {
      throw error;
    }
  };

  // handle button submission
  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    try {
      handleInputEdgeCases();
      const links = await handleTitleToLinks();
      setisLoading(true);
      setOutputState(false);
      await HandlePostAPI(links);
      await handleBackendPolling();
    } catch (error) {
      showToast(error + "", "error");
      setisLoading(false);
      setOutputState(false);
      return;
    }
  };

  const handleFromAutoComplete = (data: string) => {
    setFormValue({
      FROM: data,
      TO: formValue.TO,
    });
  };

  const handleToAutoComplete = (data: string) => {
    setFormValue({
      FROM: formValue.FROM,
      TO: data,
    });
  };

  const handleFromFocus = () => {
    setIsFromAutocompleteOpen(true);
  };

  const handleFromBlur = () => {
    setTimeout(() => {
      setIsFromAutocompleteOpen(false);
    }, 200);
  };

  const handleToFocus = () => {
    setIsToAutocompleteOpen(true);
  };

  const handleToBlur = () => {
    setTimeout(() => {
      setIsToAutocompleteOpen(false);
    }, 200);
  };
  return (
    <>
      <form action="submit" onSubmit={handleSubmit}>
        <div className="flex flex-col w-full justify-center gap-y-5 gap-x-20">
          <div className="flex flex-col justify-center" data-aos="fade-up">
            <label htmlFor="FROM">FROM</label>
            <input
              type="text"
              id="FROM"
              className="h-10 bg-gray-100/40 border-solid rounded-md"
              placeholder="  wikipedia title"
              value={formValue.FROM}
              onChange={handleChange}
              onFocus={handleFromFocus}
              onBlur={handleFromBlur}
            />
            {isFromAutocompleteOpen && (
              <Autocomplete
                data={formValue.FROM}
                setData={handleFromAutoComplete}
              />
            )}
          </div>
          <div className="flex flex-col" data-aos="fade-up">
            <label htmlFor="TO">TO</label>
            <input
              type="text"
              id="TO"
              className="h-10 bg-gray-100/40 border-solid rounded-md"
              placeholder="  wikipedia title"
              value={formValue.TO}
              onChange={handleChange}
              onFocus={handleToFocus}
              onBlur={handleToBlur}
            />
            {isToAutocompleteOpen && (
              <Autocomplete
                data={formValue.TO}
                setData={handleToAutoComplete}
              />
            )}
          </div>
        </div>
        <div className="flex justify-center items-center w-full my-12" data-aos="fade-left">
          <button 
            className='transform transition duration-300 hover:scale-110 w-[200px] bg-[#221465] rounded-md font-medium my-6 mx-auto py-4 text-white' style={{boxShadow: 'inset 0 0 10px rgba(201, 191, 255, 0.5)'}}
            disabled={isLoading}
          >
            FIND
          </button>
        </div>
      </form>
      <div>
        <LoadingBar isLoading={isLoading} />
      </div>
    </>
  );
};

export default EntryWiki;
