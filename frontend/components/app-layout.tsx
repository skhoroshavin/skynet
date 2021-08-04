import React from "react";
import {NavBar} from "./navbar/navbar";

type AppLayoutProps = {
    children: React.ReactFragment
};

export const AppLayout = ({ children }: AppLayoutProps) => {
    return <>
        <header className="bg-primary-600 text-grey-900 fixed inset-x-0 top-0 h-16 z-10">
            <NavBar/>
        </header>
        <main className="container mx-auto mt-16 p-4 flex items-start">
            {children}
        </main>
    </>
}
