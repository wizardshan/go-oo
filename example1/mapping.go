package main

// entity
//go:generate go run github.com/wizardshan/structmapper -toName Item -fromName Item -toPath ./domain -fromPath ./repository/entity  -toVar dom -fromVar ent -funcPosition from
//go:generate go run github.com/wizardshan/structmapper -toName Items -fromName Items -toPath ./domain -fromPath ./repository/entity  -toVar dom -fromVar ent -funcPosition from

// response
//go:generate go run github.com/wizardshan/structmapper -toName Item -fromName Item -toPath ./response  -fromPath ./domain -toVar resp -fromVar dom
//go:generate go run github.com/wizardshan/structmapper -toName Items -fromName Items -toPath ./response  -fromPath ./domain -toVar resp -fromVar dom
