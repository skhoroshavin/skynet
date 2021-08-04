import React from "react";
import Link from "next/link";
import {NavBrand} from "./brand";
import {NavSearch} from "./search";

type NavLinkProps = {
    text: string,
    href: string
    className?: string
}

const NavLink = ({ text, href, className }: NavLinkProps) => {
    return <Link href={href} passHref>
        <a className={`h-full p-2 flex items-center hover:bg-primary-700 font-bold transition ${className}`}>
            {text}
        </a>
    </Link>
}

const NavUser = () => {
    return <NavLink className="ml-auto px-6" text="Sign In" href="/signin"/>
}

export const NavBar = () => {
    return <nav className="h-full mx-auto container flex items-center">
        <NavBrand/>
        <NavSearch/>
        <NavUser/>
    </nav>
}
