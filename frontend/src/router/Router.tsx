
 

import {VFC, useState} from "react"
import {Route, Routes,  Link} from "react-router-dom"
import MarkdownEditor from "../components/page/Markdown"
import { Search } from "../components/page/Search"
import { Home } from "../components/page/Home"
import { MarkdownProvider } from "../hooks/providers/useMarkdownsProvider"

//MemDockの様にする？
export const Router = () => {
    return (
        <MarkdownProvider>
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/markdowns" element={<Home />} />
                <Route path="/markdowns/:id" element={<Home />} />
            </Routes>
        </MarkdownProvider>
    )
}