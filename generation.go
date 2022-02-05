package main

import (
   "bufio"
   "os"
   log "github.com/sirupsen/logrus"
)

func generateDockerfile() {
   var lcBaseContainer = "archlinux"
   var installer = "pacman"
   var installArgs = "-Syy"
   var justInstall = "--noconfirm"
   var deleteCache = "-Scc"
   const lbNonRoot = true
   const lcUser = "base"
   const lcWorkDir = "/home/" + lcUser
   var apps = make([]string, 2) // Write an append function!!! make([]string, len(s), (cap(s)+1)*2)

   apps[0] = "vim"
   apps[1] = "tmux"
   // apps[2] = "texlive-most"

   log.Debug("1")
   lcFile, lcErr := os.Create("Dockerfile")
   if lcErr != nil {
      log.Fatal(lcErr)
   }

   log.Debug("2")
   lcBufferWriter := bufio.NewWriter(lcFile)

   log.Debug("3")
   _, lcErr = lcBufferWriter.WriteString("FROM " + lcBaseContainer + "\n")
   if lcErr != nil {
      log.Fatal(lcErr)
   }

   log.Debug("4")
   _, lcErr = lcBufferWriter.WriteString("RUN " + installer + " " +
      justInstall + " " + installArgs + "\\\n")
   if lcErr != nil {
      log.Fatal(lcErr)
   }

   log.Debug("5")
   for l, lcApp := range apps {
      _, lcErr = lcBufferWriter.WriteString("\t" + lcApp + " \\\n")
      log.Debug(l)
      if lcErr != nil {
         log.Fatal(lcErr)
      }
   }

   log.Debug("6")
   _, lcErr = lcBufferWriter.WriteString("\t&& " + installer + " " + deleteCache)
   if lcErr != nil {
      log.Fatal(lcErr)
   }

   log.Debug("7")
   if lbNonRoot {
      log.Debug("8")
      _, lcErr = lcBufferWriter.WriteString(" \\\n")
      if lcErr != nil {
         log.Fatal(lcErr)
      }
      log.Debug("9")
      _, lcErr = lcBufferWriter.WriteString("\t&& useradd -m " + lcUser)
      if lcErr != nil {
         log.Fatal(lcErr)
      }
      log.Debug("10")
      _, lcErr = lcBufferWriter.WriteString("\n\nWORKDIR " + lcWorkDir)
      if lcErr != nil {
         log.Fatal(lcErr)
      }
   }
   log.Debug("11")
   lcBufferWriter.Flush()
   log.Debug("12")
}
