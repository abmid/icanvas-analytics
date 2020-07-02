/*
 * File Created: Wednesday, 1st July 2020 3:09:39 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 * 
 * Copyright (c) 2020 Author
 */

export function middleware(store, router) {
    router.beforeEach((to, from, next) => {
        const requiresAuth = to.matched.some(record => record.meta.requiredAuth);
        const currentUser = store.state.auth.currentUser;
        const appName = "iCanvas Analytics"
        const metaTitle = to.meta.title
        document.title = appName + " - " + metaTitle
        if(requiresAuth && !currentUser) {
            // if route required auth and user not authenticated
            // redirect to / or login
            next('/');
        } else if(to.path == '/' && currentUser) {        
            // if route to / or login and user already authenticated
            next("/home")     
        } else {                  
            next()
        }
    });
}