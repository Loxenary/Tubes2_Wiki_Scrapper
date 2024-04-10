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

  // Save Data for loading logic
  const [isLoading, setisLoading] = useState(false);

  // Save Data for Links to related wiki pages
  const [linksValue, setLinksValue] = useState({
    FROM: "",
    TO: "",
  });
  // Context that saves the output data
  const { setOutputData }: IOutputContext = useOutputContext();

  // Save the data of the algorithm that is used
  const { Algorithm }: SearchWikiInterface = useWikiSearchContext();

  const [isFromAutocompleteOpen, setIsFromAutocompleteOpen] = useState(false);
  const [isToAutocompleteOpen, setIsToAutocompleteOpen] = useState(false);

  // Context for turn on the visibility of the output
  const SearchContext = useContext(BoolOutputSetup);
  if (!SearchContext) {
    showToast("Context not found", "error");
    return null;
  }
  const { setOutputState } = SearchContext;

  //
  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { id, value } = event.target;
    //TODO: handle autocomplete from wikipedia api
    setFormValue((prevState) => ({
      ...prevState,
      [id]: value,
    }));
  };

  function handleOutputData(data: any) {
    const { checkcount, listPath, numpassed, time } = data;
    setOutputData(checkcount, numpassed, time, listPath);
  }

  function handleInputData() {
    const dataUsed = {
      ...formValue,
      algorithm: Algorithm,
    };
    setFormValue(dataUsed);
  }

  function handleInputEdgeCases() {
    if (formValue.FROM === formValue.TO) {
      throw new Error("The Data of From and To are the Same");
    } else if (formValue.FROM.length === 0 || formValue.TO.length === 0) {
      throw new Error("Please fill all the input fields");
    }
  }

  const HandleAPI = async () => {
    const res = await fetch("/api/postData", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(linksValue),
    });

    if (!res.ok) {
      throw new Error("failed to fetch");
    }
    const data = await res.json();
    return data;
  };

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    //TODO: implement api request to handle backend
    try {
      handleInputEdgeCases();
      await handleLinks();
      handleInputData();
      console.log(formValue);
      setisLoading(true);
      setOutputState(false);

      const data = await HandleAPI();
      setisLoading(false);

      handleOutputData(data);
      setOutputState(true);
    } catch (error) {
      showToast(error + "", "error");
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

  function removeSpace(data: string) {
    return data.replace(/\s+/g, "_");
  }

  function converLink(data: string) {
    return `https://en.wikipedia.org/wiki/${removeSpace(data)}`;
  }

  const handleLinks = async () => {
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
      setLinksValue({
        FROM: converLink(formattedFrom),
        TO: converLink(formattedTo),
      });
    } catch (error) {
      throw error;
    }
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
