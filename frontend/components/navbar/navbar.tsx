import React from "react";
import Link from "next/link";
import {NavSearch} from "./search";
import {NuclearBlast} from "../common/nuclear-blast";
import {NavUser} from "./user";

export const NavBar = () => {
    return <nav className="h-full mx-auto container flex items-center">
        <Link href={"/"} passHref>
            <a><NuclearBlast className="ml-4 py-3 h-16"/></a>
        </Link>
        <NavSearch/>
        <NavUser/>
    </nav>
}
