type CardProps = {
    children: React.ReactFragment,
    className?: string
}

export const Card = ({ className, children }: CardProps) => {
    return <div className={`p-4 bg-white rounded-lg shadow-md ${className}`}>
        {children}
    </div>
}

type CardDividerProps = {
    className?: string
}

export const CardDivider = ({ className }: CardDividerProps ) => {
    return <hr className={`border-grey-300 ${className}`}/>
}
