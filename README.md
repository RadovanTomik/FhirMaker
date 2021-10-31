
Credit goes to the samply development community since most of the code is reused from their project: https://github.com/samply/bbmri-fhir-ig 
###### V0.1
A basic go application, works only for MMCI style XML files, the application transforms
every file in the `input` folder. Output files are .json transactions.

###### V0.3
Changed the code to convert CRC cohort data model files to FHIR, output files were tested and are queriable by the SL
