# hs

The hotspot analysis tool _hs_ analyses a git repository and reports a list of files ordered by
their commit frequency and complexity:

```bash
hs -base -url=https://github.com/spring-projects/spring-data-jpa.git

Commits File                                    Lines   Complexity
109     SimpleJpaRepository.java                456     6.47
102     UserRepository.java                     256     4.12
88      QueryUtils.java                         404     7.21
60      JpaQueryMethod.java                     215     6.25
60      AbstractJpaQuery.java                   241     8.95
58      JpaRepositoryFactory.java               166     6.14
56      JpaQueryCreator.java                    218     11.30
55      PartTreeJpaQuery.java                   210     9.14
50      JpaQueryExecution.java                  184     8.00
48      JpaMetamodelEntityInformation.java      248     8.71
45      JpaQueryLookupStrategy.java             111     9.37
44      StringQuery.java                        428     11.02
44      ParameterBinder.java                    42      6.10
39      SimpleJpaQuery.java                     37      7.14
38      AbstractStringBasedJpaQuery.java        69      6.38
37      JpaRepositoryFactoryBean.java           65      4.31
37      JpaRepositoryConfigExtension.java       153     6.01
```

The idea of the tool is to spot refactoring candidates based on their complexity and change
frequency. For instance, `SimpleJpaRepository` is the most frequently changing file that has a fair
amount of complexity. But `JpaQueryCreator`'s complexity is twice as high, thus it might be
worthwhile to examine this file first.

## Usage
The tool could be run either on a local or remote git repository depending on the URL given as
command-line parameter:

```
# run hs on local repository
hs -url file:///Users/ralf/tmp/spring-data-jpa

# run hs on remote repository
hs -url https://github.com/spring-projects/spring-data-jpa.git
```

Remote repository URLs are cloned or pulled prior the analysis into _tmp_ of the users home
directory. Run `hs --help` for further usage details.
