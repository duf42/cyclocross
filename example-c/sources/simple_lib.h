
#ifndef SIMPLELIB__
#define SIMPLELIB__

/* Structure definitions */
typedef struct {
    double target;
    double current;
} Inputs_T;

typedef struct {
    double command;
} Outputs_T;


typedef struct {
    double Kp;
    double Ki;
    double Ts;
} Parameters_T;

typedef struct {
    double x;
} States_T;

/* Interface variable declaration */
extern Inputs_T Model_U;
extern Outputs_T Model_Y;
extern Parameters_T Model_P;
extern States_T Model_DW;

/* Function declaration */
extern void initialize();
extern void step();
extern void terminate();

#endif