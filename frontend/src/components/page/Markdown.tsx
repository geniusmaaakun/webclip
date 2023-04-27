// import React, { useState } from "react";
// import SimpleMde from "react-simplemde-editor";
// import "easymde/dist/easymde.min.css";

//基本的なエディタの実装
// export const MarkdownEditor = () => {
//   const [markdownValue, setMarkdownValue] = useState("Initial value");

//   const onChange = (value) => {
//     setMarkdownValue(value);
//   };

//   return <SimpleMde value={markdownValue} onChange={onChange} />;
// };

// export default MarkdownEditor;


// import React, { useState } from "react";
// import SimpleMde from "react-simplemde-editor";
// import "easymde/dist/easymde.min.css";
// import {marked} from "marked";
// import DOMPurify from "dompurify";
// import highlightjs from "highlight.js";
// import "highlight.js/styles/github.css";

// //プレビュー機能の追加
// const MarkdownEditor = () => {
//     const [markdownValue, setMarkdownValue] = useState("Initial value");
   
//     const onChange = (value) => {
//       setMarkdownValue(value);
//     };
   
//     return (
//     <>
//        <SimpleMde value={markdownValue} onChange={onChange} />
//        <div>
//           <div dangerouslySetInnerHTML={{__html: DOMPurify.sanitize(marked(markdownValue))}}></div>
//        </div>
//     </>
//     );
//    };
   
// export default MarkdownEditor;


import React, { useState } from "react";
import SimpleMde from "react-simplemde-editor";
import "easymde/dist/easymde.min.css";
import {marked} from "marked";
//npm install --save @types/dompurifyをしないといけない
import DOMPurify from "dompurify";
import highlightjs from "highlight.js";
import "highlight.js/styles/github.css";
//ハイライトをつけよう


/**
 ```js
const [test, setTest] = useState();
```
```go
sl := 1
```
 */
const MarkdownEditor = () => {
    // ハイライトの設定
  marked.setOptions({
    highlight: (code, lang) => {
      return highlightjs.highlightAuto(code, [lang]).value;
    },
  });
 
  const [markdownValue, setMarkdownValue] = useState("");
 
  const onChange = (value: string) => {
    setMarkdownValue(value);
    //保存する処理
  };
 
  return (
    <div>
      <SimpleMde value={markdownValue} onChange={onChange} />
      <div
        dangerouslySetInnerHTML={{
          __html: DOMPurify.sanitize(marked(markdownValue)),
        }}
      ></div>
    </div>
  );
 };
 
 export default MarkdownEditor;
 
