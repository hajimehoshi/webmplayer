//go:build ignore

/* The contents of this file was automatically generated by
 * dump_mode_arm_ne10.c with arguments: 48000 960
 * It contains static definitions for some pre-defined modes. */
#include <NE10_types.h>

#ifndef NE10_FFT_PARAMS48000_960
#define NE10_FFT_PARAMS48000_960
static const ne10_int32_t ne10_factors_480[64] = {
4, 40, 4, 30, 2, 15, 5, 3, 3, 1, 1, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, };
static const ne10_int32_t ne10_factors_240[64] = {
3, 20, 4, 15, 5, 3, 3, 1, 1, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, };
static const ne10_int32_t ne10_factors_120[64] = {
3, 10, 2, 15, 5, 3, 3, 1, 1, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, };
static const ne10_int32_t ne10_factors_60[64] = {
2, 5, 5, 3, 3, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, };
static const ne10_fft_cpx_float32_t ne10_twiddles_480[480] = {
{1.0000000f,0.0000000f}, {1.0000000f,-0.0000000f}, {1.0000000f,-0.0000000f},
{1.0000000f,-0.0000000f}, {0.91354543f,-0.40673664f}, {0.66913056f,-0.74314487f},
{1.0000000f,-0.0000000f}, {0.66913056f,-0.74314487f}, {-0.10452851f,-0.99452192f},
{1.0000000f,-0.0000000f}, {0.30901697f,-0.95105654f}, {-0.80901700f,-0.58778518f},
{1.0000000f,-0.0000000f}, {-0.10452851f,-0.99452192f}, {-0.97814757f,0.20791179f},
{1.0000000f,-0.0000000f}, {0.97814763f,-0.20791170f}, {0.91354543f,-0.40673664f},
{0.80901700f,-0.58778524f}, {0.66913056f,-0.74314487f}, {0.49999997f,-0.86602545f},
{0.30901697f,-0.95105654f}, {0.10452842f,-0.99452192f}, {-0.10452851f,-0.99452192f},
{-0.30901703f,-0.95105648f}, {-0.50000006f,-0.86602533f}, {-0.66913068f,-0.74314475f},
{-0.80901700f,-0.58778518f}, {-0.91354549f,-0.40673658f}, {-0.97814763f,-0.20791161f},
{1.0000000f,-0.0000000f}, {0.99862951f,-0.052335959f}, {0.99452192f,-0.10452846f},
{0.98768836f,-0.15643448f}, {0.97814763f,-0.20791170f}, {0.96592581f,-0.25881904f},
{0.95105648f,-0.30901700f}, {0.93358040f,-0.35836795f}, {0.91354543f,-0.40673664f},
{0.89100653f,-0.45399052f}, {0.86602545f,-0.50000000f}, {0.83867055f,-0.54463905f},
{0.80901700f,-0.58778524f}, {0.77714598f,-0.62932038f}, {0.74314475f,-0.66913062f},
{0.70710677f,-0.70710683f}, {0.66913056f,-0.74314487f}, {0.62932038f,-0.77714598f},
{0.58778524f,-0.80901700f}, {0.54463899f,-0.83867055f}, {0.49999997f,-0.86602545f},
{0.45399052f,-0.89100653f}, {0.40673661f,-0.91354549f}, {0.35836786f,-0.93358046f},
{0.30901697f,-0.95105654f}, {0.25881907f,-0.96592581f}, {0.20791166f,-0.97814763f},
{0.15643437f,-0.98768836f}, {0.10452842f,-0.99452192f}, {0.052335974f,-0.99862951f},
{1.0000000f,-0.0000000f}, {0.99452192f,-0.10452846f}, {0.97814763f,-0.20791170f},
{0.95105648f,-0.30901700f}, {0.91354543f,-0.40673664f}, {0.86602545f,-0.50000000f},
{0.80901700f,-0.58778524f}, {0.74314475f,-0.66913062f}, {0.66913056f,-0.74314487f},
{0.58778524f,-0.80901700f}, {0.49999997f,-0.86602545f}, {0.40673661f,-0.91354549f},
{0.30901697f,-0.95105654f}, {0.20791166f,-0.97814763f}, {0.10452842f,-0.99452192f},
{-4.3711388e-08f,-1.0000000f}, {-0.10452851f,-0.99452192f}, {-0.20791174f,-0.97814757f},
{-0.30901703f,-0.95105648f}, {-0.40673670f,-0.91354543f}, {-0.50000006f,-0.86602533f},
{-0.58778518f,-0.80901700f}, {-0.66913068f,-0.74314475f}, {-0.74314493f,-0.66913044f},
{-0.80901700f,-0.58778518f}, {-0.86602539f,-0.50000006f}, {-0.91354549f,-0.40673658f},
{-0.95105654f,-0.30901679f}, {-0.97814763f,-0.20791161f}, {-0.99452192f,-0.10452849f},
{1.0000000f,-0.0000000f}, {0.98768836f,-0.15643448f}, {0.95105648f,-0.30901700f},
{0.89100653f,-0.45399052f}, {0.80901700f,-0.58778524f}, {0.70710677f,-0.70710683f},
{0.58778524f,-0.80901700f}, {0.45399052f,-0.89100653f}, {0.30901697f,-0.95105654f},
{0.15643437f,-0.98768836f}, {-4.3711388e-08f,-1.0000000f}, {-0.15643445f,-0.98768836f},
{-0.30901703f,-0.95105648f}, {-0.45399061f,-0.89100647f}, {-0.58778518f,-0.80901700f},
{-0.70710677f,-0.70710677f}, {-0.80901700f,-0.58778518f}, {-0.89100659f,-0.45399037f},
{-0.95105654f,-0.30901679f}, {-0.98768836f,-0.15643445f}, {-1.0000000f,8.7422777e-08f},
{-0.98768830f,0.15643461f}, {-0.95105654f,0.30901697f}, {-0.89100653f,0.45399055f},
{-0.80901694f,0.58778536f}, {-0.70710665f,0.70710689f}, {-0.58778507f,0.80901712f},
{-0.45399022f,0.89100665f}, {-0.30901709f,0.95105648f}, {-0.15643452f,0.98768830f},
{1.0000000f,-0.0000000f}, {0.99991435f,-0.013089596f}, {0.99965733f,-0.026176950f},
{0.99922901f,-0.039259817f}, {0.99862951f,-0.052335959f}, {0.99785894f,-0.065403134f},
{0.99691731f,-0.078459099f}, {0.99580491f,-0.091501623f}, {0.99452192f,-0.10452846f},
{0.99306846f,-0.11753740f}, {0.99144489f,-0.13052620f}, {0.98965138f,-0.14349262f},
{0.98768836f,-0.15643448f}, {0.98555607f,-0.16934951f}, {0.98325491f,-0.18223552f},
{0.98078525f,-0.19509032f}, {0.97814763f,-0.20791170f}, {0.97534233f,-0.22069745f},
{0.97236991f,-0.23344538f}, {0.96923089f,-0.24615330f}, {0.96592581f,-0.25881904f},
{0.96245521f,-0.27144045f}, {0.95881975f,-0.28401536f}, {0.95501995f,-0.29654160f},
{0.95105648f,-0.30901700f}, {0.94693011f,-0.32143945f}, {0.94264150f,-0.33380687f},
{0.93819129f,-0.34611708f}, {0.93358040f,-0.35836795f}, {0.92880952f,-0.37055743f},
{0.92387956f,-0.38268346f}, {0.91879117f,-0.39474389f}, {0.91354543f,-0.40673664f},
{0.90814316f,-0.41865975f}, {0.90258527f,-0.43051112f}, {0.89687270f,-0.44228873f},
{0.89100653f,-0.45399052f}, {0.88498765f,-0.46561453f}, {0.87881708f,-0.47715878f},
{0.87249601f,-0.48862126f}, {0.86602545f,-0.50000000f}, {0.85940641f,-0.51129311f},
{0.85264015f,-0.52249855f}, {0.84572786f,-0.53361452f}, {0.83867055f,-0.54463905f},
{0.83146960f,-0.55557024f}, {0.82412618f,-0.56640625f}, {0.81664151f,-0.57714522f},
{0.80901700f,-0.58778524f}, {0.80125380f,-0.59832460f}, {0.79335332f,-0.60876143f},
{0.78531694f,-0.61909395f}, {0.77714598f,-0.62932038f}, {0.76884180f,-0.63943899f},
{0.76040596f,-0.64944810f}, {0.75183982f,-0.65934587f}, {0.74314475f,-0.66913062f},
{0.73432249f,-0.67880076f}, {0.72537434f,-0.68835455f}, {0.71630192f,-0.69779050f},
{0.70710677f,-0.70710683f}, {0.69779044f,-0.71630198f}, {0.68835455f,-0.72537440f},
{0.67880070f,-0.73432255f}, {0.66913056f,-0.74314487f}, {0.65934581f,-0.75183982f},
{0.64944804f,-0.76040596f}, {0.63943899f,-0.76884186f}, {0.62932038f,-0.77714598f},
{0.61909395f,-0.78531694f}, {0.60876137f,-0.79335338f}, {0.59832460f,-0.80125386f},
{0.58778524f,-0.80901700f}, {0.57714516f,-0.81664151f}, {0.56640625f,-0.82412618f},
{0.55557019f,-0.83146960f}, {0.54463899f,-0.83867055f}, {0.53361452f,-0.84572786f},
{0.52249849f,-0.85264015f}, {0.51129311f,-0.85940641f}, {0.49999997f,-0.86602545f},
{0.48862118f,-0.87249601f}, {0.47715876f,-0.87881708f}, {0.46561447f,-0.88498765f},
{0.45399052f,-0.89100653f}, {0.44228867f,-0.89687276f}, {0.43051103f,-0.90258533f},
{0.41865975f,-0.90814316f}, {0.40673661f,-0.91354549f}, {0.39474380f,-0.91879129f},
{0.38268343f,-0.92387956f}, {0.37055740f,-0.92880958f}, {0.35836786f,-0.93358046f},
{0.34611705f,-0.93819135f}, {0.33380681f,-0.94264150f}, {0.32143947f,-0.94693011f},
{0.30901697f,-0.95105654f}, {0.29654151f,-0.95501995f}, {0.28401533f,-0.95881975f},
{0.27144039f,-0.96245527f}, {0.25881907f,-0.96592581f}, {0.24615327f,-0.96923089f},
{0.23344530f,-0.97236991f}, {0.22069745f,-0.97534233f}, {0.20791166f,-0.97814763f},
{0.19509023f,-0.98078531f}, {0.18223552f,-0.98325491f}, {0.16934945f,-0.98555607f},
{0.15643437f,-0.98768836f}, {0.14349259f,-0.98965138f}, {0.13052613f,-0.99144489f},
{0.11753740f,-0.99306846f}, {0.10452842f,-0.99452192f}, {0.091501534f,-0.99580491f},
{0.078459084f,-0.99691731f}, {0.065403074f,-0.99785894f}, {0.052335974f,-0.99862951f},
{0.039259788f,-0.99922901f}, {0.026176875f,-0.99965733f}, {0.013089597f,-0.99991435f},
{1.0000000f,-0.0000000f}, {0.99965733f,-0.026176950f}, {0.99862951f,-0.052335959f},
{0.99691731f,-0.078459099f}, {0.99452192f,-0.10452846f}, {0.99144489f,-0.13052620f},
{0.98768836f,-0.15643448f}, {0.98325491f,-0.18223552f}, {0.97814763f,-0.20791170f},
{0.97236991f,-0.23344538f}, {0.96592581f,-0.25881904f}, {0.95881975f,-0.28401536f},
{0.95105648f,-0.30901700f}, {0.94264150f,-0.33380687f}, {0.93358040f,-0.35836795f},
{0.92387956f,-0.38268346f}, {0.91354543f,-0.40673664f}, {0.90258527f,-0.43051112f},
{0.89100653f,-0.45399052f}, {0.87881708f,-0.47715878f}, {0.86602545f,-0.50000000f},
{0.85264015f,-0.52249855f}, {0.83867055f,-0.54463905f}, {0.82412618f,-0.56640625f},
{0.80901700f,-0.58778524f}, {0.79335332f,-0.60876143f}, {0.77714598f,-0.62932038f},
{0.76040596f,-0.64944810f}, {0.74314475f,-0.66913062f}, {0.72537434f,-0.68835455f},
{0.70710677f,-0.70710683f}, {0.68835455f,-0.72537440f}, {0.66913056f,-0.74314487f},
{0.64944804f,-0.76040596f}, {0.62932038f,-0.77714598f}, {0.60876137f,-0.79335338f},
{0.58778524f,-0.80901700f}, {0.56640625f,-0.82412618f}, {0.54463899f,-0.83867055f},
{0.52249849f,-0.85264015f}, {0.49999997f,-0.86602545f}, {0.47715876f,-0.87881708f},
{0.45399052f,-0.89100653f}, {0.43051103f,-0.90258533f}, {0.40673661f,-0.91354549f},
{0.38268343f,-0.92387956f}, {0.35836786f,-0.93358046f}, {0.33380681f,-0.94264150f},
{0.30901697f,-0.95105654f}, {0.28401533f,-0.95881975f}, {0.25881907f,-0.96592581f},
{0.23344530f,-0.97236991f}, {0.20791166f,-0.97814763f}, {0.18223552f,-0.98325491f},
{0.15643437f,-0.98768836f}, {0.13052613f,-0.99144489f}, {0.10452842f,-0.99452192f},
{0.078459084f,-0.99691731f}, {0.052335974f,-0.99862951f}, {0.026176875f,-0.99965733f},
{-4.3711388e-08f,-1.0000000f}, {-0.026176963f,-0.99965733f}, {-0.052336060f,-0.99862951f},
{-0.078459173f,-0.99691731f}, {-0.10452851f,-0.99452192f}, {-0.13052621f,-0.99144489f},
{-0.15643445f,-0.98768836f}, {-0.18223560f,-0.98325491f}, {-0.20791174f,-0.97814757f},
{-0.23344538f,-0.97236991f}, {-0.25881916f,-0.96592581f}, {-0.28401542f,-0.95881969f},
{-0.30901703f,-0.95105648f}, {-0.33380687f,-0.94264150f}, {-0.35836795f,-0.93358040f},
{-0.38268352f,-0.92387950f}, {-0.40673670f,-0.91354543f}, {-0.43051112f,-0.90258527f},
{-0.45399061f,-0.89100647f}, {-0.47715873f,-0.87881708f}, {-0.50000006f,-0.86602533f},
{-0.52249867f,-0.85264009f}, {-0.54463905f,-0.83867055f}, {-0.56640631f,-0.82412612f},
{-0.58778518f,-0.80901700f}, {-0.60876143f,-0.79335332f}, {-0.62932050f,-0.77714586f},
{-0.64944804f,-0.76040596f}, {-0.66913068f,-0.74314475f}, {-0.68835467f,-0.72537428f},
{-0.70710677f,-0.70710677f}, {-0.72537446f,-0.68835449f}, {-0.74314493f,-0.66913044f},
{-0.76040596f,-0.64944804f}, {-0.77714604f,-0.62932026f}, {-0.79335332f,-0.60876143f},
{-0.80901700f,-0.58778518f}, {-0.82412624f,-0.56640613f}, {-0.83867055f,-0.54463899f},
{-0.85264021f,-0.52249849f}, {-0.86602539f,-0.50000006f}, {-0.87881714f,-0.47715873f},
{-0.89100659f,-0.45399037f}, {-0.90258527f,-0.43051112f}, {-0.91354549f,-0.40673658f},
{-0.92387956f,-0.38268328f}, {-0.93358040f,-0.35836792f}, {-0.94264150f,-0.33380675f},
{-0.95105654f,-0.30901679f}, {-0.95881975f,-0.28401530f}, {-0.96592587f,-0.25881892f},
{-0.97236991f,-0.23344538f}, {-0.97814763f,-0.20791161f}, {-0.98325491f,-0.18223536f},
{-0.98768836f,-0.15643445f}, {-0.99144489f,-0.13052608f}, {-0.99452192f,-0.10452849f},
{-0.99691737f,-0.078459039f}, {-0.99862957f,-0.052335810f}, {-0.99965733f,-0.026176952f},
{1.0000000f,-0.0000000f}, {0.99922901f,-0.039259817f}, {0.99691731f,-0.078459099f},
{0.99306846f,-0.11753740f}, {0.98768836f,-0.15643448f}, {0.98078525f,-0.19509032f},
{0.97236991f,-0.23344538f}, {0.96245521f,-0.27144045f}, {0.95105648f,-0.30901700f},
{0.93819129f,-0.34611708f}, {0.92387956f,-0.38268346f}, {0.90814316f,-0.41865975f},
{0.89100653f,-0.45399052f}, {0.87249601f,-0.48862126f}, {0.85264015f,-0.52249855f},
{0.83146960f,-0.55557024f}, {0.80901700f,-0.58778524f}, {0.78531694f,-0.61909395f},
{0.76040596f,-0.64944810f}, {0.73432249f,-0.67880076f}, {0.70710677f,-0.70710683f},
{0.67880070f,-0.73432255f}, {0.64944804f,-0.76040596f}, {0.61909395f,-0.78531694f},
{0.58778524f,-0.80901700f}, {0.55557019f,-0.83146960f}, {0.52249849f,-0.85264015f},
{0.48862118f,-0.87249601f}, {0.45399052f,-0.89100653f}, {0.41865975f,-0.90814316f},
{0.38268343f,-0.92387956f}, {0.34611705f,-0.93819135f}, {0.30901697f,-0.95105654f},
{0.27144039f,-0.96245527f}, {0.23344530f,-0.97236991f}, {0.19509023f,-0.98078531f},
{0.15643437f,-0.98768836f}, {0.11753740f,-0.99306846f}, {0.078459084f,-0.99691731f},
{0.039259788f,-0.99922901f}, {-4.3711388e-08f,-1.0000000f}, {-0.039259877f,-0.99922901f},
{-0.078459173f,-0.99691731f}, {-0.11753749f,-0.99306846f}, {-0.15643445f,-0.98768836f},
{-0.19509032f,-0.98078525f}, {-0.23344538f,-0.97236991f}, {-0.27144048f,-0.96245521f},
{-0.30901703f,-0.95105648f}, {-0.34611711f,-0.93819129f}, {-0.38268352f,-0.92387950f},
{-0.41865984f,-0.90814310f}, {-0.45399061f,-0.89100647f}, {-0.48862135f,-0.87249595f},
{-0.52249867f,-0.85264009f}, {-0.55557036f,-0.83146954f}, {-0.58778518f,-0.80901700f},
{-0.61909389f,-0.78531694f}, {-0.64944804f,-0.76040596f}, {-0.67880076f,-0.73432249f},
{-0.70710677f,-0.70710677f}, {-0.73432249f,-0.67880070f}, {-0.76040596f,-0.64944804f},
{-0.78531694f,-0.61909389f}, {-0.80901700f,-0.58778518f}, {-0.83146966f,-0.55557019f},
{-0.85264021f,-0.52249849f}, {-0.87249607f,-0.48862115f}, {-0.89100659f,-0.45399037f},
{-0.90814322f,-0.41865960f}, {-0.92387956f,-0.38268328f}, {-0.93819135f,-0.34611690f},
{-0.95105654f,-0.30901679f}, {-0.96245521f,-0.27144048f}, {-0.97236991f,-0.23344538f},
{-0.98078531f,-0.19509031f}, {-0.98768836f,-0.15643445f}, {-0.99306846f,-0.11753736f},
{-0.99691737f,-0.078459039f}, {-0.99922901f,-0.039259743f}, {-1.0000000f,8.7422777e-08f},
{-0.99922901f,0.039259918f}, {-0.99691731f,0.078459218f}, {-0.99306846f,0.11753753f},
{-0.98768830f,0.15643461f}, {-0.98078525f,0.19509049f}, {-0.97236985f,0.23344554f},
{-0.96245515f,0.27144065f}, {-0.95105654f,0.30901697f}, {-0.93819135f,0.34611705f},
{-0.92387956f,0.38268346f}, {-0.90814316f,0.41865975f}, {-0.89100653f,0.45399055f},
{-0.87249601f,0.48862129f}, {-0.85264015f,0.52249861f}, {-0.83146960f,0.55557030f},
{-0.80901694f,0.58778536f}, {-0.78531688f,0.61909401f}, {-0.76040590f,0.64944816f},
{-0.73432243f,0.67880082f}, {-0.70710665f,0.70710689f}, {-0.67880058f,0.73432261f},
{-0.64944792f,0.76040608f}, {-0.61909378f,0.78531706f}, {-0.58778507f,0.80901712f},
{-0.55557001f,0.83146977f}, {-0.52249837f,0.85264033f}, {-0.48862100f,0.87249613f},
{-0.45399022f,0.89100665f}, {-0.41865945f,0.90814328f}, {-0.38268313f,0.92387968f},
{-0.34611672f,0.93819147f}, {-0.30901709f,0.95105648f}, {-0.27144054f,0.96245521f},
{-0.23344545f,0.97236991f}, {-0.19509038f,0.98078525f}, {-0.15643452f,0.98768830f},
{-0.11753743f,0.99306846f}, {-0.078459114f,0.99691731f}, {-0.039259821f,0.99922901f},
};
static const ne10_fft_cpx_float32_t ne10_twiddles_240[240] = {
{1.0000000f,0.0000000f}, {1.0000000f,-0.0000000f}, {1.0000000f,-0.0000000f},
{1.0000000f,-0.0000000f}, {0.91354543f,-0.40673664f}, {0.66913056f,-0.74314487f},
{1.0000000f,-0.0000000f}, {0.66913056f,-0.74314487f}, {-0.10452851f,-0.99452192f},
{1.0000000f,-0.0000000f}, {0.30901697f,-0.95105654f}, {-0.80901700f,-0.58778518f},
{1.0000000f,-0.0000000f}, {-0.10452851f,-0.99452192f}, {-0.97814757f,0.20791179f},
{1.0000000f,-0.0000000f}, {0.99452192f,-0.10452846f}, {0.97814763f,-0.20791170f},
{0.95105648f,-0.30901700f}, {0.91354543f,-0.40673664f}, {0.86602545f,-0.50000000f},
{0.80901700f,-0.58778524f}, {0.74314475f,-0.66913062f}, {0.66913056f,-0.74314487f},
{0.58778524f,-0.80901700f}, {0.49999997f,-0.86602545f}, {0.40673661f,-0.91354549f},
{0.30901697f,-0.95105654f}, {0.20791166f,-0.97814763f}, {0.10452842f,-0.99452192f},
{1.0000000f,-0.0000000f}, {0.97814763f,-0.20791170f}, {0.91354543f,-0.40673664f},
{0.80901700f,-0.58778524f}, {0.66913056f,-0.74314487f}, {0.49999997f,-0.86602545f},
{0.30901697f,-0.95105654f}, {0.10452842f,-0.99452192f}, {-0.10452851f,-0.99452192f},
{-0.30901703f,-0.95105648f}, {-0.50000006f,-0.86602533f}, {-0.66913068f,-0.74314475f},
{-0.80901700f,-0.58778518f}, {-0.91354549f,-0.40673658f}, {-0.97814763f,-0.20791161f},
{1.0000000f,-0.0000000f}, {0.95105648f,-0.30901700f}, {0.80901700f,-0.58778524f},
{0.58778524f,-0.80901700f}, {0.30901697f,-0.95105654f}, {-4.3711388e-08f,-1.0000000f},
{-0.30901703f,-0.95105648f}, {-0.58778518f,-0.80901700f}, {-0.80901700f,-0.58778518f},
{-0.95105654f,-0.30901679f}, {-1.0000000f,8.7422777e-08f}, {-0.95105654f,0.30901697f},
{-0.80901694f,0.58778536f}, {-0.58778507f,0.80901712f}, {-0.30901709f,0.95105648f},
{1.0000000f,-0.0000000f}, {0.99965733f,-0.026176950f}, {0.99862951f,-0.052335959f},
{0.99691731f,-0.078459099f}, {0.99452192f,-0.10452846f}, {0.99144489f,-0.13052620f},
{0.98768836f,-0.15643448f}, {0.98325491f,-0.18223552f}, {0.97814763f,-0.20791170f},
{0.97236991f,-0.23344538f}, {0.96592581f,-0.25881904f}, {0.95881975f,-0.28401536f},
{0.95105648f,-0.30901700f}, {0.94264150f,-0.33380687f}, {0.93358040f,-0.35836795f},
{0.92387956f,-0.38268346f}, {0.91354543f,-0.40673664f}, {0.90258527f,-0.43051112f},
{0.89100653f,-0.45399052f}, {0.87881708f,-0.47715878f}, {0.86602545f,-0.50000000f},
{0.85264015f,-0.52249855f}, {0.83867055f,-0.54463905f}, {0.82412618f,-0.56640625f},
{0.80901700f,-0.58778524f}, {0.79335332f,-0.60876143f}, {0.77714598f,-0.62932038f},
{0.76040596f,-0.64944810f}, {0.74314475f,-0.66913062f}, {0.72537434f,-0.68835455f},
{0.70710677f,-0.70710683f}, {0.68835455f,-0.72537440f}, {0.66913056f,-0.74314487f},
{0.64944804f,-0.76040596f}, {0.62932038f,-0.77714598f}, {0.60876137f,-0.79335338f},
{0.58778524f,-0.80901700f}, {0.56640625f,-0.82412618f}, {0.54463899f,-0.83867055f},
{0.52249849f,-0.85264015f}, {0.49999997f,-0.86602545f}, {0.47715876f,-0.87881708f},
{0.45399052f,-0.89100653f}, {0.43051103f,-0.90258533f}, {0.40673661f,-0.91354549f},
{0.38268343f,-0.92387956f}, {0.35836786f,-0.93358046f}, {0.33380681f,-0.94264150f},
{0.30901697f,-0.95105654f}, {0.28401533f,-0.95881975f}, {0.25881907f,-0.96592581f},
{0.23344530f,-0.97236991f}, {0.20791166f,-0.97814763f}, {0.18223552f,-0.98325491f},
{0.15643437f,-0.98768836f}, {0.13052613f,-0.99144489f}, {0.10452842f,-0.99452192f},
{0.078459084f,-0.99691731f}, {0.052335974f,-0.99862951f}, {0.026176875f,-0.99965733f},
{1.0000000f,-0.0000000f}, {0.99862951f,-0.052335959f}, {0.99452192f,-0.10452846f},
{0.98768836f,-0.15643448f}, {0.97814763f,-0.20791170f}, {0.96592581f,-0.25881904f},
{0.95105648f,-0.30901700f}, {0.93358040f,-0.35836795f}, {0.91354543f,-0.40673664f},
{0.89100653f,-0.45399052f}, {0.86602545f,-0.50000000f}, {0.83867055f,-0.54463905f},
{0.80901700f,-0.58778524f}, {0.77714598f,-0.62932038f}, {0.74314475f,-0.66913062f},
{0.70710677f,-0.70710683f}, {0.66913056f,-0.74314487f}, {0.62932038f,-0.77714598f},
{0.58778524f,-0.80901700f}, {0.54463899f,-0.83867055f}, {0.49999997f,-0.86602545f},
{0.45399052f,-0.89100653f}, {0.40673661f,-0.91354549f}, {0.35836786f,-0.93358046f},
{0.30901697f,-0.95105654f}, {0.25881907f,-0.96592581f}, {0.20791166f,-0.97814763f},
{0.15643437f,-0.98768836f}, {0.10452842f,-0.99452192f}, {0.052335974f,-0.99862951f},
{-4.3711388e-08f,-1.0000000f}, {-0.052336060f,-0.99862951f}, {-0.10452851f,-0.99452192f},
{-0.15643445f,-0.98768836f}, {-0.20791174f,-0.97814757f}, {-0.25881916f,-0.96592581f},
{-0.30901703f,-0.95105648f}, {-0.35836795f,-0.93358040f}, {-0.40673670f,-0.91354543f},
{-0.45399061f,-0.89100647f}, {-0.50000006f,-0.86602533f}, {-0.54463905f,-0.83867055f},
{-0.58778518f,-0.80901700f}, {-0.62932050f,-0.77714586f}, {-0.66913068f,-0.74314475f},
{-0.70710677f,-0.70710677f}, {-0.74314493f,-0.66913044f}, {-0.77714604f,-0.62932026f},
{-0.80901700f,-0.58778518f}, {-0.83867055f,-0.54463899f}, {-0.86602539f,-0.50000006f},
{-0.89100659f,-0.45399037f}, {-0.91354549f,-0.40673658f}, {-0.93358040f,-0.35836792f},
{-0.95105654f,-0.30901679f}, {-0.96592587f,-0.25881892f}, {-0.97814763f,-0.20791161f},
{-0.98768836f,-0.15643445f}, {-0.99452192f,-0.10452849f}, {-0.99862957f,-0.052335810f},
{1.0000000f,-0.0000000f}, {0.99691731f,-0.078459099f}, {0.98768836f,-0.15643448f},
{0.97236991f,-0.23344538f}, {0.95105648f,-0.30901700f}, {0.92387956f,-0.38268346f},
{0.89100653f,-0.45399052f}, {0.85264015f,-0.52249855f}, {0.80901700f,-0.58778524f},
{0.76040596f,-0.64944810f}, {0.70710677f,-0.70710683f}, {0.64944804f,-0.76040596f},
{0.58778524f,-0.80901700f}, {0.52249849f,-0.85264015f}, {0.45399052f,-0.89100653f},
{0.38268343f,-0.92387956f}, {0.30901697f,-0.95105654f}, {0.23344530f,-0.97236991f},
{0.15643437f,-0.98768836f}, {0.078459084f,-0.99691731f}, {-4.3711388e-08f,-1.0000000f},
{-0.078459173f,-0.99691731f}, {-0.15643445f,-0.98768836f}, {-0.23344538f,-0.97236991f},
{-0.30901703f,-0.95105648f}, {-0.38268352f,-0.92387950f}, {-0.45399061f,-0.89100647f},
{-0.52249867f,-0.85264009f}, {-0.58778518f,-0.80901700f}, {-0.64944804f,-0.76040596f},
{-0.70710677f,-0.70710677f}, {-0.76040596f,-0.64944804f}, {-0.80901700f,-0.58778518f},
{-0.85264021f,-0.52249849f}, {-0.89100659f,-0.45399037f}, {-0.92387956f,-0.38268328f},
{-0.95105654f,-0.30901679f}, {-0.97236991f,-0.23344538f}, {-0.98768836f,-0.15643445f},
{-0.99691737f,-0.078459039f}, {-1.0000000f,8.7422777e-08f}, {-0.99691731f,0.078459218f},
{-0.98768830f,0.15643461f}, {-0.97236985f,0.23344554f}, {-0.95105654f,0.30901697f},
{-0.92387956f,0.38268346f}, {-0.89100653f,0.45399055f}, {-0.85264015f,0.52249861f},
{-0.80901694f,0.58778536f}, {-0.76040590f,0.64944816f}, {-0.70710665f,0.70710689f},
{-0.64944792f,0.76040608f}, {-0.58778507f,0.80901712f}, {-0.52249837f,0.85264033f},
{-0.45399022f,0.89100665f}, {-0.38268313f,0.92387968f}, {-0.30901709f,0.95105648f},
{-0.23344545f,0.97236991f}, {-0.15643452f,0.98768830f}, {-0.078459114f,0.99691731f},
};
static const ne10_fft_cpx_float32_t ne10_twiddles_120[120] = {
{1.0000000f,0.0000000f}, {1.0000000f,-0.0000000f}, {1.0000000f,-0.0000000f},
{1.0000000f,-0.0000000f}, {0.91354543f,-0.40673664f}, {0.66913056f,-0.74314487f},
{1.0000000f,-0.0000000f}, {0.66913056f,-0.74314487f}, {-0.10452851f,-0.99452192f},
{1.0000000f,-0.0000000f}, {0.30901697f,-0.95105654f}, {-0.80901700f,-0.58778518f},
{1.0000000f,-0.0000000f}, {-0.10452851f,-0.99452192f}, {-0.97814757f,0.20791179f},
{1.0000000f,-0.0000000f}, {0.97814763f,-0.20791170f}, {0.91354543f,-0.40673664f},
{0.80901700f,-0.58778524f}, {0.66913056f,-0.74314487f}, {0.49999997f,-0.86602545f},
{0.30901697f,-0.95105654f}, {0.10452842f,-0.99452192f}, {-0.10452851f,-0.99452192f},
{-0.30901703f,-0.95105648f}, {-0.50000006f,-0.86602533f}, {-0.66913068f,-0.74314475f},
{-0.80901700f,-0.58778518f}, {-0.91354549f,-0.40673658f}, {-0.97814763f,-0.20791161f},
{1.0000000f,-0.0000000f}, {0.99862951f,-0.052335959f}, {0.99452192f,-0.10452846f},
{0.98768836f,-0.15643448f}, {0.97814763f,-0.20791170f}, {0.96592581f,-0.25881904f},
{0.95105648f,-0.30901700f}, {0.93358040f,-0.35836795f}, {0.91354543f,-0.40673664f},
{0.89100653f,-0.45399052f}, {0.86602545f,-0.50000000f}, {0.83867055f,-0.54463905f},
{0.80901700f,-0.58778524f}, {0.77714598f,-0.62932038f}, {0.74314475f,-0.66913062f},
{0.70710677f,-0.70710683f}, {0.66913056f,-0.74314487f}, {0.62932038f,-0.77714598f},
{0.58778524f,-0.80901700f}, {0.54463899f,-0.83867055f}, {0.49999997f,-0.86602545f},
{0.45399052f,-0.89100653f}, {0.40673661f,-0.91354549f}, {0.35836786f,-0.93358046f},
{0.30901697f,-0.95105654f}, {0.25881907f,-0.96592581f}, {0.20791166f,-0.97814763f},
{0.15643437f,-0.98768836f}, {0.10452842f,-0.99452192f}, {0.052335974f,-0.99862951f},
{1.0000000f,-0.0000000f}, {0.99452192f,-0.10452846f}, {0.97814763f,-0.20791170f},
{0.95105648f,-0.30901700f}, {0.91354543f,-0.40673664f}, {0.86602545f,-0.50000000f},
{0.80901700f,-0.58778524f}, {0.74314475f,-0.66913062f}, {0.66913056f,-0.74314487f},
{0.58778524f,-0.80901700f}, {0.49999997f,-0.86602545f}, {0.40673661f,-0.91354549f},
{0.30901697f,-0.95105654f}, {0.20791166f,-0.97814763f}, {0.10452842f,-0.99452192f},
{-4.3711388e-08f,-1.0000000f}, {-0.10452851f,-0.99452192f}, {-0.20791174f,-0.97814757f},
{-0.30901703f,-0.95105648f}, {-0.40673670f,-0.91354543f}, {-0.50000006f,-0.86602533f},
{-0.58778518f,-0.80901700f}, {-0.66913068f,-0.74314475f}, {-0.74314493f,-0.66913044f},
{-0.80901700f,-0.58778518f}, {-0.86602539f,-0.50000006f}, {-0.91354549f,-0.40673658f},
{-0.95105654f,-0.30901679f}, {-0.97814763f,-0.20791161f}, {-0.99452192f,-0.10452849f},
{1.0000000f,-0.0000000f}, {0.98768836f,-0.15643448f}, {0.95105648f,-0.30901700f},
{0.89100653f,-0.45399052f}, {0.80901700f,-0.58778524f}, {0.70710677f,-0.70710683f},
{0.58778524f,-0.80901700f}, {0.45399052f,-0.89100653f}, {0.30901697f,-0.95105654f},
{0.15643437f,-0.98768836f}, {-4.3711388e-08f,-1.0000000f}, {-0.15643445f,-0.98768836f},
{-0.30901703f,-0.95105648f}, {-0.45399061f,-0.89100647f}, {-0.58778518f,-0.80901700f},
{-0.70710677f,-0.70710677f}, {-0.80901700f,-0.58778518f}, {-0.89100659f,-0.45399037f},
{-0.95105654f,-0.30901679f}, {-0.98768836f,-0.15643445f}, {-1.0000000f,8.7422777e-08f},
{-0.98768830f,0.15643461f}, {-0.95105654f,0.30901697f}, {-0.89100653f,0.45399055f},
{-0.80901694f,0.58778536f}, {-0.70710665f,0.70710689f}, {-0.58778507f,0.80901712f},
{-0.45399022f,0.89100665f}, {-0.30901709f,0.95105648f}, {-0.15643452f,0.98768830f},
};
static const ne10_fft_cpx_float32_t ne10_twiddles_60[60] = {
{1.0000000f,0.0000000f}, {1.0000000f,-0.0000000f}, {1.0000000f,-0.0000000f},
{1.0000000f,-0.0000000f}, {0.91354543f,-0.40673664f}, {0.66913056f,-0.74314487f},
{1.0000000f,-0.0000000f}, {0.66913056f,-0.74314487f}, {-0.10452851f,-0.99452192f},
{1.0000000f,-0.0000000f}, {0.30901697f,-0.95105654f}, {-0.80901700f,-0.58778518f},
{1.0000000f,-0.0000000f}, {-0.10452851f,-0.99452192f}, {-0.97814757f,0.20791179f},
{1.0000000f,-0.0000000f}, {0.99452192f,-0.10452846f}, {0.97814763f,-0.20791170f},
{0.95105648f,-0.30901700f}, {0.91354543f,-0.40673664f}, {0.86602545f,-0.50000000f},
{0.80901700f,-0.58778524f}, {0.74314475f,-0.66913062f}, {0.66913056f,-0.74314487f},
{0.58778524f,-0.80901700f}, {0.49999997f,-0.86602545f}, {0.40673661f,-0.91354549f},
{0.30901697f,-0.95105654f}, {0.20791166f,-0.97814763f}, {0.10452842f,-0.99452192f},
{1.0000000f,-0.0000000f}, {0.97814763f,-0.20791170f}, {0.91354543f,-0.40673664f},
{0.80901700f,-0.58778524f}, {0.66913056f,-0.74314487f}, {0.49999997f,-0.86602545f},
{0.30901697f,-0.95105654f}, {0.10452842f,-0.99452192f}, {-0.10452851f,-0.99452192f},
{-0.30901703f,-0.95105648f}, {-0.50000006f,-0.86602533f}, {-0.66913068f,-0.74314475f},
{-0.80901700f,-0.58778518f}, {-0.91354549f,-0.40673658f}, {-0.97814763f,-0.20791161f},
{1.0000000f,-0.0000000f}, {0.95105648f,-0.30901700f}, {0.80901700f,-0.58778524f},
{0.58778524f,-0.80901700f}, {0.30901697f,-0.95105654f}, {-4.3711388e-08f,-1.0000000f},
{-0.30901703f,-0.95105648f}, {-0.58778518f,-0.80901700f}, {-0.80901700f,-0.58778518f},
{-0.95105654f,-0.30901679f}, {-1.0000000f,8.7422777e-08f}, {-0.95105654f,0.30901697f},
{-0.80901694f,0.58778536f}, {-0.58778507f,0.80901712f}, {-0.30901709f,0.95105648f},
};
static const ne10_fft_state_float32_t ne10_fft_state_float32_t_480 = {
120,
(ne10_int32_t *)ne10_factors_480,
(ne10_fft_cpx_float32_t *)ne10_twiddles_480,
NULL,
(ne10_fft_cpx_float32_t *)&ne10_twiddles_480[120],
/* is_forward_scaled = true */
(ne10_int32_t) 1,
/* is_backward_scaled = false */
(ne10_int32_t) 0,
};
static const arch_fft_state cfg_arch_480 = {
1,
(void *)&ne10_fft_state_float32_t_480,
};

static const ne10_fft_state_float32_t ne10_fft_state_float32_t_240 = {
60,
(ne10_int32_t *)ne10_factors_240,
(ne10_fft_cpx_float32_t *)ne10_twiddles_240,
NULL,
(ne10_fft_cpx_float32_t *)&ne10_twiddles_240[60],
/* is_forward_scaled = true */
(ne10_int32_t) 1,
/* is_backward_scaled = false */
(ne10_int32_t) 0,
};
static const arch_fft_state cfg_arch_240 = {
1,
(void *)&ne10_fft_state_float32_t_240,
};

static const ne10_fft_state_float32_t ne10_fft_state_float32_t_120 = {
30,
(ne10_int32_t *)ne10_factors_120,
(ne10_fft_cpx_float32_t *)ne10_twiddles_120,
NULL,
(ne10_fft_cpx_float32_t *)&ne10_twiddles_120[30],
/* is_forward_scaled = true */
(ne10_int32_t) 1,
/* is_backward_scaled = false */
(ne10_int32_t) 0,
};
static const arch_fft_state cfg_arch_120 = {
1,
(void *)&ne10_fft_state_float32_t_120,
};

static const ne10_fft_state_float32_t ne10_fft_state_float32_t_60 = {
15,
(ne10_int32_t *)ne10_factors_60,
(ne10_fft_cpx_float32_t *)ne10_twiddles_60,
NULL,
(ne10_fft_cpx_float32_t *)&ne10_twiddles_60[15],
/* is_forward_scaled = true */
(ne10_int32_t) 1,
/* is_backward_scaled = false */
(ne10_int32_t) 0,
};
static const arch_fft_state cfg_arch_60 = {
1,
(void *)&ne10_fft_state_float32_t_60,
};

#endif  /* end NE10_FFT_PARAMS48000_960 */
