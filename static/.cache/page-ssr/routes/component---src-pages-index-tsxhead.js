"use strict";
exports.id = "component---src-pages-index-tsxhead";
exports.ids = ["component---src-pages-index-tsxhead"];
exports.modules = {

/***/ "./src/pages/index.tsx?export=head":
/*!*****************************************!*\
  !*** ./src/pages/index.tsx?export=head ***!
  \*****************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   Head: () => (/* binding */ Head),
/* harmony export */   "default": () => (__WEBPACK_DEFAULT_EXPORT__)
/* harmony export */ });
/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react */ "react");
/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_0__);

const pageStyles = {
  color: "#232129",
  padding: 96,
  fontFamily: "-apple-system, Roboto, sans-serif, serif"
};
const headingStyles = {
  marginTop: 0,
  marginBottom: 64,
  maxWidth: 320
};
const headingAccentStyles = {
  color: "#663399"
};
const paragraphStyles = {
  marginBottom: 48
};
const codeStyles = {
  color: "#8A6534",
  padding: 4,
  backgroundColor: "#FFF4DB",
  fontSize: "1.25rem",
  borderRadius: 4
};
const listStyles = {
  marginBottom: 96,
  paddingLeft: 0
};
const doclistStyles = {
  paddingLeft: 0
};
const listItemStyles = {
  fontWeight: 300,
  fontSize: 24,
  maxWidth: 560,
  marginBottom: 30
};
const linkStyle = {
  color: "#8954A8",
  fontWeight: "bold",
  fontSize: 16,
  verticalAlign: "5%"
};
const docLinkStyle = {
  ...linkStyle,
  listStyleType: "none",
  display: `inline-block`,
  marginBottom: 24,
  marginRight: 12
};
const descriptionStyle = {
  color: "#232129",
  fontSize: 14,
  marginTop: 10,
  marginBottom: 0,
  lineHeight: 1.25
};
const buttonStyle = {
  border: "1px solid #088413"
};
const badgeStyle = {
  color: "#fff",
  backgroundColor: "#088413",
  border: "1px solid #088413",
  fontSize: 11,
  fontWeight: "bold",
  letterSpacing: 1,
  borderRadius: 4,
  padding: "4px 6px",
  display: "inline-block",
  position: "relative",
  top: -2,
  marginLeft: 10,
  lineHeight: 1
};
const links = [{
  text: "Tutorial",
  url: "https://www.gatsbyjs.com/docs/tutorial/getting-started/",
  description: "A great place to get started if you're new to web development. Designed to guide you through setting up your first Gatsby site.",
  color: "#E95800"
}, {
  text: "How to Guides",
  url: "https://www.gatsbyjs.com/docs/how-to/",
  description: "Practical step-by-step guides to help you achieve a specific goal. Most useful when you're trying to get something done.",
  color: "#1099A8"
}, {
  text: "Reference Guides",
  url: "https://www.gatsbyjs.com/docs/reference/",
  description: "Nitty-gritty technical descriptions of how Gatsby works. Most useful when you need detailed information about Gatsby's APIs.",
  color: "#BC027F"
}, {
  text: "Conceptual Guides",
  url: "https://www.gatsbyjs.com/docs/conceptual/",
  description: "Big-picture explanations of higher-level Gatsby concepts. Most useful for building understanding of a particular topic.",
  color: "#0D96F2"
}, {
  text: "Plugin Library",
  url: "https://www.gatsbyjs.com/plugins",
  description: "Add functionality and customize your Gatsby site or app with thousands of plugins built by our amazing developer community.",
  color: "#8EB814"
}, {
  text: "Build and Host",
  url: "https://www.gatsbyjs.com/cloud",
  badge: true,
  description: "Now you’re ready to show the world! Give your Gatsby site superpowers: Build and host on Gatsby Cloud. Get started for free!",
  color: "#663399"
}];
const IndexPage = () => {
  return /*#__PURE__*/react__WEBPACK_IMPORTED_MODULE_0__.createElement("main", {
    style: pageStyles
  }, /*#__PURE__*/react__WEBPACK_IMPORTED_MODULE_0__.createElement("h1", {
    style: headingStyles
  }, "Welcome", /*#__PURE__*/react__WEBPACK_IMPORTED_MODULE_0__.createElement("br", null), /*#__PURE__*/react__WEBPACK_IMPORTED_MODULE_0__.createElement("span", {
    style: headingAccentStyles
  }, "\u2014 to the Copilot Demo Zone \uD83C\uDF89\uD83C\uDF89\uD83C\uDF89")), /*#__PURE__*/react__WEBPACK_IMPORTED_MODULE_0__.createElement("p", {
    style: paragraphStyles
  }, "Enter some data in the form:"), /*#__PURE__*/react__WEBPACK_IMPORTED_MODULE_0__.createElement("form", null, /*#__PURE__*/react__WEBPACK_IMPORTED_MODULE_0__.createElement("label", null, "ID", /*#__PURE__*/react__WEBPACK_IMPORTED_MODULE_0__.createElement("input", {
    type: "text",
    id: "id",
    name: "id"
  })), /*#__PURE__*/react__WEBPACK_IMPORTED_MODULE_0__.createElement("br", null), /*#__PURE__*/react__WEBPACK_IMPORTED_MODULE_0__.createElement("label", null, "Chewiness (how long this request will take to process)", /*#__PURE__*/react__WEBPACK_IMPORTED_MODULE_0__.createElement("input", {
    style: buttonStyle,
    type: "text",
    id: "chewiness",
    name: "chewiness"
  })), /*#__PURE__*/react__WEBPACK_IMPORTED_MODULE_0__.createElement("br", null), /*#__PURE__*/react__WEBPACK_IMPORTED_MODULE_0__.createElement("label", null, "Data (string only, please)", /*#__PURE__*/react__WEBPACK_IMPORTED_MODULE_0__.createElement("input", {
    style: buttonStyle,
    type: "text",
    id: "data",
    name: "data"
  })), /*#__PURE__*/react__WEBPACK_IMPORTED_MODULE_0__.createElement("br", null), /*#__PURE__*/react__WEBPACK_IMPORTED_MODULE_0__.createElement("input", {
    type: "submit",
    style: buttonStyle,
    id: "submit",
    name: "submit",
    value: "submit"
  })), /*#__PURE__*/react__WEBPACK_IMPORTED_MODULE_0__.createElement("img", {
    alt: "Gatsby G Logo",
    src: "data:image/svg+xml,%3Csvg width='24' height='24' fill='none' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M12 2a10 10 0 110 20 10 10 0 010-20zm0 2c-3.73 0-6.86 2.55-7.75 6L14 19.75c3.45-.89 6-4.02 6-7.75h-5.25v1.5h3.45a6.37 6.37 0 01-3.89 4.44L6.06 9.69C7 7.31 9.3 5.63 12 5.63c2.13 0 4 1.04 5.18 2.65l1.23-1.06A7.959 7.959 0 0012 4zm-8 8a8 8 0 008 8c.04 0 .09 0-8-8z' fill='%23639'/%3E%3C/svg%3E"
  }));
};
/* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = (IndexPage);
const Head = () => /*#__PURE__*/react__WEBPACK_IMPORTED_MODULE_0__.createElement("title", null, "Home Page");

/***/ })

};
;
//# sourceMappingURL=component---src-pages-index-tsxhead.js.map