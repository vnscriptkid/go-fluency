## Requirements
- svc must be able to answer synchronously under 200ms (SLA)

## High-level design

## Questions
- What are basic info sent to fraud-detection-svc?
    - sender acc, sender past txn ...
    - receiver acc
    - ip,location,time
- Where does data used to train model come from?
    - custoner svc data
    - actual fraud cases in the past