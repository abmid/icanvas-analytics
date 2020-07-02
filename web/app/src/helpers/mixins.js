var iMixin = {
    methods: {
      $$_TOAST_SHOW(status = "", title, message){
        var addClass = ""
        switch (status) {
            case "success":
                addClass = "bg-success"
                break;
            case "danger":
                addClass = "bg-danger"
                break;
            case "warning":
                addClass = "bg-warning"
                break;
            default:
                break;
        }
        $(document).Toasts('create', {
            class: addClass, 
            icon: 'fas fa-envelope fa-lg',
            position: 'topRight',
            title: title,
            autohide: true,
            delay: 2350,
            body: message
        }) 
      },              
    }
}

export default iMixin