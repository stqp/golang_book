#! bin/bash

dir_list=(
"01_tutorial"
"02_program_structure"
"03_basic_data_types"
"04_composite_types"
"05_functions"
"06_methods"
"07_interfaces"
"08_goroutines_and_channels"
"09_concurrency_with_shared_varianbles"
"10_packages_and_go_tool"
"11_testing"
"12_reflection"
"13_low_level_programming"
)

exercise_num=(
"12"
"5"
"12"
"14"
"19"
"5"
"18"
"18"
"15"
"4"
"7"
"12"
"4"
)

for i in ${!dir_list[@]}
do
  ch_dir=${dir_list[$i]}
  mkdir $ch_dir 
  cd $ch_dir

  for j in `seq ${exercise_num[$i]}`
  do
    ex_dir=""
    if [ $j -gt 9 ]
    then
      ex_dir="ex"$j
    else
      ex_dir="ex0"$j
    fi
    mkdir $ex_dir
    cd $ex_dir
    touch "main.go"
    cd "../"
  done
  cd "../"
done

