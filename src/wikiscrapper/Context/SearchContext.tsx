import { createContext, useContext, useState } from "react";

export interface SearchWikiInterface {
  fromValue: string;
  toValue: string;
  setData: (from: string, to: string) => void;
}

const SearchContext = createContext<SearchWikiInterface | undefined>(undefined);
// Create a context with default values undefined

export const useWikiSearchContext = () => {
  const context = useContext(SearchContext);
  if (!context) {
    throw new Error(
      "useWikiSearchContext must be used within a WikiSearchContextProvider"
    );
  }
  return context;
};

interface SearchWikiContextProvider {
  children: React.ReactNode;
}

export const WikiSearchContextProvider: React.FC<SearchWikiContextProvider> = ({
  children,
}) => {
  const [fromValue, setFromValue] = useState("");
  const [toValue, setToValue] = useState("");

  const setData = (from: string, to: string) => {
    setFromValue(from);
    setToValue(to);
  };

  return (
    <SearchContext.Provider value={{ fromValue, toValue, setData }}>
      {children}
    </SearchContext.Provider>
  );
};
