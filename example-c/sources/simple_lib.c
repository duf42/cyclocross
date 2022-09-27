
#include "simple_lib.h"

/* Global variables definition */
Inputs_T Model_U;
Outputs_T Model_Y;
Parameters_T Model_P;
States_T Model_DW;

void initialize(){
    /* Inputs */
    Model_U.current = 0.0;
    Model_U.target  = 0.0;
    /* Outputs */
    Model_Y.command = 0.0;
    /* Parameters */
    Model_P.Kp      = 0.8;
    Model_P.Ki      = 0.2;
    Model_P.Ts      = 0.1;
    /* States */
    Model_DW.x      = 0;
}

void step(){

    double error;

    error = Model_U.target - Model_U.current;

    Model_DW.x = error * Model_P.Ki * Model_P.Ts + Model_DW.x;
    Model_Y.command = error * Model_P.Kp + Model_DW.x;

}

void terminate(){
    
}