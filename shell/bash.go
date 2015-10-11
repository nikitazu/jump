package shell

var Bash = Shell(`# Put the line below in ~/.bashrc or ~/bash_profile:
#
#   eval "$(jump shell bash)"	
#
# The following lines are autogenerated:

__jump_prompt_command() {
  jump chdir
}

[[ "$PROMPT_COMMAND" =~ __jump_prompt_command ]] || {
  PROMPT_COMMAND="__jump_prompt_command;$PROMPT_COMMAND"
}

j() {
  local dir="$(jump cd $@)"
  test -d "$dir"  && cd "$dir"
}

complete -o filenames -C 'jump hint "${COMP_LINE/#j /}" --smart' j
`)
