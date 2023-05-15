//import { FaSave } from "react-icons/fa";

//保存ボタンサンプル

export const  SaveButton = () => {
    const onSave =() => {
        console.log("save");
    }
    return (
        /*
      <button onClick={props.onClick}>
        <FaSave />
      </button>
      */

      <button onClick={onSave}>
        Save
      </button>
    );
  }