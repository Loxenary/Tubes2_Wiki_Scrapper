import React from "react";

import { ToastContainer, toast } from "react-toastify";
import "react-toastify/ReactToastify.css";

const Toast = () => {
  return (
    <ToastContainer
      position="top-center"
      autoClose={2000}
      hideProgressBar={false}
      newestOnTop={false}
      stacked={true}
      closeOnClick
      rtl={false}
      pauseOnFocusLoss
      draggable
    />
  );
};

export const showToast = (message: string, type: string) => {
  switch (type) {
    case "success":
      toast.success(message);
      break;
    case "error":
      toast.error(message);
      break;
    case "info":
      toast.info(message);
      break;
    case "warning":
      toast.warn(message);
      break;
    default:
      toast.error(message);
      break;
  }
};

export default Toast;
