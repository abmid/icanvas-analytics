/*
 * File Created: Monday, 29th June 2020 11:57:34 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */


export function getTemplateCurrent(){
    const template = localStorage.getItem("template");

    if(!template) {
        return null;
    }

    return template;    
}

export function getUser() {
    const userStr = localStorage.getItem("icanvas_user");

    if(!userStr) {
        return null;
    }

    return JSON.parse(userStr);

}