"use client";
import { showToast } from "@/components/toast";
import { WikipediaExistChecker } from "@/pages/api/wiki";
import {
  SearchWikiInterface,
  useWikiSearchContext,
} from "@/Context/SearchContext";
import { useContext, useState, ChangeEvent, FormEvent } from "react";
import { BoolOutputSetup } from "./page";
import { IOutputContext } from "../Output/outputData";
import { useOutputContext } from "@/Context/OutputContext";
import { LoadingBar } from "@/components/loading";
import Autocomplete from "./autocomplete";
const EntryWiki = () => {
  // Save Data for Input Textarea
  const [formValue, setFormValue] = useState({
    FROM: "",
    TO: "",
  });

  // Save Data for loading animation logic
  const [isLoading, setisLoading] = useState(false);

  // Context that saves the output data, used in output page later
  const { setOutputData }: IOutputContext = useOutputContext();

  // Save the data of the algorithm that is used
  const { Algorithm }: SearchWikiInterface = useWikiSearchContext();

  // Used to control the state of the autocomplete when user click on the input field
  const [isFromAutocompleteOpen, setIsFromAutocompleteOpen] = useState(false);
  const [isToAutocompleteOpen, setIsToAutocompleteOpen] = useState(false);

  // Context for turn on the visibility of the output
  const SearchContext = useContext(BoolOutputSetup);
  if (!SearchContext) {
    showToast("Context not found", "error");
    return null;
  }
  const { setOutputState } = SearchContext;

  // used to save the state of current input field
  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { id, value } = event.target;
    //TODO: handle autocomplete from wikipedia api
    setFormValue((prevState) => ({
      ...prevState,
      [id]: value,
    }));
  };

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
  const HandleAPI = async (data: any) => {
    console.log("Links: " + JSON.stringify(data));
    console.log("Title : " + JSON.stringify(formValue));
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
    const output = await res.json();
    return output;
  };

  // formatting data used later in creating wikipedia links
  function removeSpace(data: string) {
    return data.replace(/\s+/g, "_");
  }

  // convert a title into a wikipedia link
  function convertLink(data: string) {
    return `https://en.wikipedia.org/wiki/${removeSpace(data)}`;
  }

  // handle any invalid links, if valid, convert the data to a link and save it into links state value
  const handleTitleToLinks = async () => {
    const { FROM, TO } = formValue;
    const formattedFrom = removeSpace(FROM);
    const formattedFromLower = FROM.toLowerCase();
    const formattedTo = removeSpace(TO);
    const formattedToLower = TO.toLowerCase();

    try {
      const responseFrom = await WikipediaExistChecker(formattedFrom);

      const responseFromLower = responseFrom.map((title) =>
        title.toLowerCase()
      );

      if (!responseFromLower.includes(formattedFromLower)) {
        throw new Error(`${FROM} title can not found on Wikipedia`);
      }

      const responseTo = await WikipediaExistChecker(formattedTo);

      const responseToLower = responseTo.map((title) => title.toLowerCase());
      if (!responseToLower.includes(formattedToLower)) {
        throw new Error(`${TO} title can not found on Wikipedia`);
      }
      const data = {
        FROM: convertLink(formattedFrom),
        TO: convertLink(formattedTo),
      };
      console.log("Links Should be : " + JSON.stringify(data));
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
      const data = await HandleAPI(links);
      handleOutputData(data);
    } catch (error) {
      showToast(error + "", "error");
      return;
    } finally {
      setOutputState(true);
      setisLoading(false);
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
        <div className="w-full justify-center flex gap-x-20">
          <div className="flex flex-col justify-center">
            <label htmlFor="FROM">FROM</label>
            <input
              type="text"
              id="FROM"
              className="bg-gray-100"
              placeholder="wikipedia title"
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
          <div className="flex flex-col">
            <label htmlFor="TO">TO</label>
            <input
              type="text"
              id="TO"
              className="bg-gray-100"
              placeholder="wikipedia title"
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
        <div className="flex justify-center items-center w-full my-12">
          <button
            className="text-white bg-blue-400 w-20 h-10 text-xl"
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
