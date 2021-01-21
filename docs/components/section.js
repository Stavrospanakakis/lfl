const Section = (props) => {
    return (
        <div className="px-2 py-3 mt-12" id={props.id}>
            <div className="font-bold text-2xl lg:text-4xl">{props.title}</div>
            <div className="text-lg lg:text-xl italic text-gray-600">{props.subtitle}</div>
            <div className="">{props.description}</div>
            {props.children}
        </div>
    )
}

export default Section