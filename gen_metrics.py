#!/usr/bin/env python

import pandas as pd
import numpy as np
import sqlite3

size = 10_000

start_time = '2022-11-03T08:17:32.11'

df = pd.DataFrame({
    'time': pd.date_range(start_time, freq='1357ms', periods=size),
    'mem': 227540992 + pd.Series(np.random.normal(scale=10793, size=size)).cumsum(),
    'cpu': 30.7 + pd.Series(np.random.normal(size=size)).cumsum(),
})
df['mem'] = df['mem'].astype(np.int64)
df['cpu'] = np.round(df['cpu'], 2)
df = df.melt(['time'], var_name='metric')
df.sort_values('time', inplace=True)
conn = sqlite3.connect('embed/metrics.db')
df.to_sql('metrics', conn, index=False)
