import { createContext, useContext, useState } from "react";

export interface SearchWikiInterface {
  Algorithm: string;
  setAlgorithm: (Algorithm: string) => void;
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
  const [Algorithm, setAlgorithm] = useState("BFS");
  return (
    <SearchContext.Provider
      value={{
        Algorithm,
        setAlgorithm,
      }}
    >
      {children}
    </SearchContext.Provider>
  );
};
