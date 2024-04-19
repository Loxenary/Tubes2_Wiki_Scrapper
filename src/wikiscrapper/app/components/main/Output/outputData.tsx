import PathInterface from "./PathData";
export interface IOutputContext {
  checkcount: string;
  numpassed: string;
  time: string;
  listPath: PathInterface[];
  setOutputData(
    checkcount: string,
    numpassed: string,
    time: string,
    pathList: PathInterface[] | undefined
  ): void;
}
