
// Copyright (c) 2020 Mr. Coxall All rights reserved
//
// Created by: Mariam Kasim
// Created on: Mar 2023
// This file contains the JS functions for index.html

'use strict'
/**
 * This function calculates area of a triangle
 */
function calculate () {
  // input
  const base = parseInt(document.getElementById('base').value)
  const height = parseInt(document.getElementById('height').value)

  // process
  const area = base * height / 2

  // output
  document.getElementById('area').innerHTML = 'Area is: ' + area + ' cm²'
}
