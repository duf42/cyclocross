
#include <stdio.h>

#include "simple_lib.h"

void main(){

    // [ SETUP ]
    int nOK, nKO;

    // [ EXERCISE ]
    initialize();

    Model_U.current = 2.0;
    Model_U.target  = 10.0;

    Model_P.Kp      = 0.5;

    step();

    // [ VERIFY ]
    nKO = 0; nOK = 0;
    if (Model_Y.command != 4.0){
        printf("[INVALID] Model_Y.command = %f (expected = %f)\n", Model_Y.command, 4.0);
        nKO++;
    } else {
        nOK++;
    }

    if (nKO > 0){
        printf("FAILED: %d verification failed\n", nKO);
    } else {
        printf("SUCCESS: %d verification success\n", nOK);
    }

}