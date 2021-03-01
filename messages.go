package main

import (
	"strings"
)

const MsgVersion = "I'm NecronicaBot v6.0.0, written in Golang!"

const MsgAddAliasTargetNotFound = "Target not found! Are you sure the target name is correct?"
const MsgAddAliasSuccess = "Alias added! **%s** is now also known as **%s**!"
const MsgAddAliasDuplicateFound = "Duplicate alias **%s** found! Please remove first with the *alias remove* command"
const MsgGetAliasSuccess = "**%s** is also known as **%s**!"
const MsgGetAliasFailed = "Sorry, I can't find an alias for **%s**..."
const MsgRemoveAliasSuccess = "Done! **%s** is not longer an alias! ^^b"
const MsgRemoveAliasFailure = "Sorry, I can't find an alias named **%s**..."
const MsgFindFoundWithAlias = "I found **%s** (aka **%s**)! ```%s```"
const MsgFindFound = "I found **%s**! ```%s```"
const MsgFindFail = "Sorry...I can't find what you are looking for >_<"

const MsgHelpQuery = "Did you do it correctly? ```%s```"
const MsgAddAliasHelp = "add-alias: Adds an alias to a 'find'\n\t> Usage: !necro add-alias <alias_name> = <target_name>\n\t(if success, you can then do '!necro, find <alias_name>')" 
const MsgRemoveAliasHelp = "remove-alias: Removes an alias\n\t> Usage: !necro remove-alias <alias_name>"
const MsgGetAliasHelp = "get-alias: Displays an alias\n\t> Usage: !necro get-alias <alias_name>"
const MsgResistHelp = "resist: Check CoC resistance!\n\t> Usage: !necro resist <active> vs <passive>"
const MsgFindHelp =  "find: Use this command to find something\n\t> Usage: !necro find <something>" 

var MsgHelp = strings.Join([]string{"```", MsgFindHelp, MsgGetAliasHelp, MsgAddAliasHelp, MsgRemoveAliasHelp, "```"}, "\n")

const MsgGenericFail = "Sorry, something went wrong...contact Momo? (´・ω・`)"
