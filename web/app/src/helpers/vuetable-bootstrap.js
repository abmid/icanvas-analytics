export default {
  table: {
    tableClass: 'table table-striped table-bordered',
    ascendingIcon: 'glyphicon glyphicon-chevron-up',
    descendingIcon: 'glyphicon glyphicon-chevron-down',
    handleIcon: 'glyphicon glyphicon-menu-hamburger',
    renderIcon: function (classes) {
      return `<span class="${classes.join(' ')}"></span>`
    }
  },
  pagination: {
    wrapperClass: "pagination float-right",
    activeClass: "btn-primary text-light",
    disabledClass: "disabled",
    pageClass: "btn btn-border",
    linkClass: "btn btn-border",
    icons: {
      first: "",
      prev: "",
      next: "",
      last: ""
    }
  },
  paginationInfo: {
    infoClass: "float-left"
  }
}