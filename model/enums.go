package model

// Apache Avro 1.10.2 Specification
// https://avro.apache.org/docs/current/spec.html

//go:generate enumeration -package=model -name=SchemaType -suffix=Enum -items=boolean,int,long,float,double,byte,string,complex,enum,ref

//go:generate enumeration -package=model -name=TemplateType -suffix=Enum -items=model,operation,namespace
//go:generate enumeration -package=model -name=LanguageExt -suffix=Enum -items=unknown,go,mod,ts,js,html,css
//go:generate enumeration -package=model -name=TemplateExt -suffix=Enum -items=unknown,gotmpl,handlebars
//go:generate enumeration -package=model -name=InputSource -suffix=Enum -items=unknown,yaml,json
