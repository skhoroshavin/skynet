import React, {FormEvent, useState} from "react";
import {SearchIcon} from "@heroicons/react/solid";

export const NavSearch = () => {
    const [searchValue, setSearchValue] = useState("")

    const handleSearch = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        console.log(searchValue)
    }

    return <form className="mx-4 h-10 rounded-md overflow-hidden" onSubmit={handleSearch}>
        <SearchIcon className="absolute m-2 h-6"/>
        <input className="p-2 pl-10 focus:outline-none
                          bg-primary-500 placeholder-grey-700
                          hover:bg-primary-300 hover:placeholder-grey-600
                          focus:bg-primary-300 focus:placeholder-grey-600"
               id="search" type="search" placeholder="Search"  onChange={(e) => setSearchValue(e.target.value)}/>
    </form>
}
