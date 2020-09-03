/*
 * File Created: Tuesday, 30th June 2020 12:09:08 am
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */


import axios from 'axios';
import config from "@/configs/config"

const service = axios.create({
  baseURL: config.apiUrl,
  timeout: 5000, // request timeout,
  withCredentials: true
});

export default service;