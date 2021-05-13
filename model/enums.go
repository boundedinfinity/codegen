package model

// Apache Avro 1.10.2 Specification
// https://avro.apache.org/docs/current/spec.html

//--go:generate enumeration -package=model -name=PrimitiveSchemaType -items=boolean,int,long,float,double,byte,string
//--go:generate enumeration -package=model -name=ComplexSchemaType -items=record,enum,array,map,fixed
//go:generate enumeration -package=model -name=SchemaType -items=boolean,int,long,float,double,byte,string,complex,enum,ref

//go:generate enumeration -package=model -name=TemplateType -items=model,operation,namespace
//go:generate enumeration -package=model -name=LanguageExt -items=unknown,go,mod,ts,js,html,css
//go:generate enumeration -package=model -name=TemplateExt -items=unknown,gotmpl,handlebars
