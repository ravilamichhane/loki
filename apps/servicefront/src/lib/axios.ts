import axios, { AxiosError, AxiosRequestConfig, AxiosResponse } from "axios";
import { ERRSTATUS, ERR_TOKEN_EXPIRED, ERR_UNAUTHORIZED, HTTP_STATUS_UNAUTHORIZED } from "./constants";

export type ErrorResponse = {
  success: false;
  error: string;
  errors?: FieldErrors;
  path?: string;
  status?: ERRSTATUS;
};

export type AjaxError = AxiosError<ErrorResponse>;

type SuccessResponse<T> = {
  success: true;
  message: string;
  data: T;
};

export type FieldErrors = Record<string, string>;
export type AjaxResponse<T> = Promise<SuccessResponse<T> | ErrorResponse>;

export const BASE_URL = "http://api.service.local/api/v1/";

const axiosInstance = axios.create({
  baseURL: BASE_URL,
});

axiosInstance.interceptors.request.use(
  async (config) => {
    config.headers["Authorization"] = `Bearer Add JWT HERE`;
    return config;
  },
  (error) => {
    Promise.reject(error);
  },
);

axiosInstance.interceptors.response.use(
  (response) => {
    return response;
  },

  async function (error: AjaxError) {
    if (error.response) {
      if (error.response.status === HTTP_STATUS_UNAUTHORIZED) {
        if (error.response.data.status === ERR_TOKEN_EXPIRED) {
          // Add refresh token logic here
        }
        if (error.response.data.status === ERR_UNAUTHORIZED) {
          // Add logout logic here
        }
      }

      return;
    }
    return Promise.reject(error);
  },
);

export function fetcher<T, U = any>(params: AxiosRequestConfig<U>) {
  return axiosInstance(params).then((res: AxiosResponse<SuccessResponse<T>, U>) => res.data);
}

type QueryParams = Record<string, string | number | boolean>;

export function Get<T, U = any>(url: string, query?: QueryParams, config?: AxiosRequestConfig): AjaxResponse<T> {
  return fetcher<T, U>({
    url,
    params: query,
    ...config,
  });
}

export function Post<T, U = any>(url: string, data?: U, config?: AxiosRequestConfig): AjaxResponse<T> {
  return fetcher<T, U>({
    url,
    method: "POST",
    data,
    ...config,
  });
}

export function Put<T, U = any>(url: string, data?: U, config?: AxiosRequestConfig): AjaxResponse<T> {
  return fetcher<T, U>({
    url,
    method: "PUT",
    data,
    ...config,
  });
}

export function Patch<T, U = any>(url: string, data?: U, config?: AxiosRequestConfig): AjaxResponse<T> {
  return fetcher<T, U>({
    url,
    method: "PATCH",
    data,
    ...config,
  });
}

export function Delete<T, U = any>(url: string, data?: U, config?: AxiosRequestConfig): AjaxResponse<T> {
  return fetcher<T, U>({
    url,
    method: "DELETE",
    data,
    ...config,
  });
}
