# Metrics specification

This document describes the metrics exposed by `saptune_exporter`.

General notes:
- All the metrics are _namespaced_ with the prefix `saptune`, which is followed by a _subsystem_, and both are in turn composed into a _Fully Qualified Name_ (FQN) of each metrics.
- All the metrics and labels _names_ are in snake_case, as conventional with Prometheus. That said, as much as we'll try to keep this consistent throughout the project, the label _values_ may not actually follow this convention, though (e.g. value is a hostname).


These are the currently implemented subsystems.

1. [Solution](#solution)
2. [Note](#note)
3. [Misc](#meta)


## Solution

0. [Sample](../test/solution.metrics)
1. [`saptune_solution_enabled`](#saptune_solution_enabled)
2. [`saptune_solution_compliant`](#saptune_solution_compliant)


### `saptune_solution_enabled`

### Description

Show which Saptune solution is enabled. A value of 1 means solution is enabled, 0 is not enabled.

#### Labels

- `solution_name`: the name of the soluton

### `saptune_solution_compliant`

### Description

Show if the given solution is compliant to saptune standards. A value of 1 meains is compliant, 0 not compliant

#### Labels

- `solution_name`: the name of the soluton

## Note

0. [Sample](../test/note.metrics)
1. [`saptune_note_enabled`](#saptune_note_enabled)

### `saptune_note_enabled`

### Description

Show which notes are enabled by ID. The value 1 means is enabled. If metric label  is not present , this means the SAP notes are not enabled.

### Label

- `note_id`: this indicate the note id of SAP note.
- `note_desc`: provide description about the note enabled.
## Meta

The meta collector collects all "meta" metrics/information about saptune. (version etc.)

0. [Sample](../test/meta.metrics)
1. [`saptune_meta_version`](#saptune_meta_version)

### Description

Show the version of saptune. The value is the major version. example `2` means saptune binary is running with 2 as major version.

