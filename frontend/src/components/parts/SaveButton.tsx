//import { FaSave } from "react-icons/fa";
import { useUpdateMarkdown } from "../../hooks/markdowns/useMarkdowns";
import {AxiosRequestConfig} from "axios";

//保存ボタンサンプル

interface Props {
  value: string;
  id: string;
}

export const  SaveButton = (props: Props) => {
    const {id, value} = props;
    const {updateMarkdown } = useUpdateMarkdown();
    const onSave =() => {
        const controller = new AbortController();
        const options: AxiosRequestConfig = {
          signal: controller.signal, //AbortControllerとAxiosの紐付け
        };
        console.log(id);
        updateMarkdown(id, value, options);
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