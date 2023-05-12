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

//ハイライトをつけよう
import React from "react";

import { Editor } from "../parts/Editor";

import {useParams} from "react-router-dom"

/**
 ```js
const [test, setTest] = useState();
```
```go
sl := 1
```
 */

// type Props = {
//   markdowns: Markdown[];
// };

export const MarkdownEditor = React.memo(() => {
  //const { markdowns } = props;

  //urlからidを取得する
  const {id} = useParams();

  if (!id) {
    return null;
  }
  return (
    <div>
      <Editor id={id}/>
    </div>
  );
 });
 
 
