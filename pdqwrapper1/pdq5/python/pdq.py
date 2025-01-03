# This file was automatically generated by SWIG (http://www.swig.org).
# Version 1.3.38
#
# Do not make changes to this file unless you know what you are doing--modify
# the SWIG interface file instead.
# This file is compatible with both classic and new-style classes.

from sys import version_info
if version_info >= (2,6,0):
    def swig_import_helper():
        from os.path import dirname
        import imp
        try:
            fp, pathname, description = imp.find_module('_pdq', [dirname(__file__)])
            _mod = imp.load_module('_pdq', fp, pathname, description)
        finally:
            if fp is not None: fp.close()
        return _mod
    _pdq = swig_import_helper()
    del swig_import_helper
else:
    import _pdq
del version_info
try:
    _swig_property = property
except NameError:
    pass # Python < 2.2 doesn't have 'property'.
def _swig_setattr_nondynamic(self,class_type,name,value,static=1):
    if (name == "thisown"): return self.this.own(value)
    if (name == "this"):
        if type(value).__name__ == 'SwigPyObject':
            self.__dict__[name] = value
            return
    method = class_type.__swig_setmethods__.get(name,None)
    if method: return method(self,value)
    if (not static) or hasattr(self,name):
        self.__dict__[name] = value
    else:
        raise AttributeError("You cannot add attributes to %s" % self)

def _swig_setattr(self,class_type,name,value):
    return _swig_setattr_nondynamic(self,class_type,name,value,0)

def _swig_getattr(self,class_type,name):
    if (name == "thisown"): return self.this.own()
    method = class_type.__swig_getmethods__.get(name,None)
    if method: return method(self)
    raise AttributeError(name)

def _swig_repr(self):
    try: strthis = "proxy of " + self.this.__repr__()
    except: strthis = ""
    return "<%s.%s; %s >" % (self.__class__.__module__, self.__class__.__name__, strthis,)

try:
    _object = object
    _newclass = 1
except AttributeError:
    class _object : pass
    _newclass = 0


TRUE = _pdq.TRUE
FALSE = _pdq.FALSE
MAXNODES = _pdq.MAXNODES
MAXBUF = _pdq.MAXBUF
MAXSTREAMS = _pdq.MAXSTREAMS
MAXCHARS = _pdq.MAXCHARS
VOID = _pdq.VOID
OPEN = _pdq.OPEN
CLOSED = _pdq.CLOSED
MEM = _pdq.MEM
CEN = _pdq.CEN
DLY = _pdq.DLY
MSQ = _pdq.MSQ
ISRV = _pdq.ISRV
FCFS = _pdq.FCFS
PSHR = _pdq.PSHR
LCFS = _pdq.LCFS
TERM = _pdq.TERM
TRANS = _pdq.TRANS
BATCH = _pdq.BATCH
EXACT = _pdq.EXACT
APPROX = _pdq.APPROX
CANON = _pdq.CANON
VISITS = _pdq.VISITS
DEMAND = _pdq.DEMAND
PDQ_SP = _pdq.PDQ_SP
PDQ_MP = _pdq.PDQ_MP
TOL = _pdq.TOL
class SYSTAT_TYPE(_object):
    __swig_setmethods__ = {}
    __setattr__ = lambda self, name, value: _swig_setattr(self, SYSTAT_TYPE, name, value)
    __swig_getmethods__ = {}
    __getattr__ = lambda self, name: _swig_getattr(self, SYSTAT_TYPE, name)
    __repr__ = _swig_repr
    __swig_setmethods__["response"] = _pdq.SYSTAT_TYPE_response_set
    __swig_getmethods__["response"] = _pdq.SYSTAT_TYPE_response_get
    if _newclass:response = _swig_property(_pdq.SYSTAT_TYPE_response_get, _pdq.SYSTAT_TYPE_response_set)
    __swig_setmethods__["thruput"] = _pdq.SYSTAT_TYPE_thruput_set
    __swig_getmethods__["thruput"] = _pdq.SYSTAT_TYPE_thruput_get
    if _newclass:thruput = _swig_property(_pdq.SYSTAT_TYPE_thruput_get, _pdq.SYSTAT_TYPE_thruput_set)
    __swig_setmethods__["residency"] = _pdq.SYSTAT_TYPE_residency_set
    __swig_getmethods__["residency"] = _pdq.SYSTAT_TYPE_residency_get
    if _newclass:residency = _swig_property(_pdq.SYSTAT_TYPE_residency_get, _pdq.SYSTAT_TYPE_residency_set)
    __swig_setmethods__["physmem"] = _pdq.SYSTAT_TYPE_physmem_set
    __swig_getmethods__["physmem"] = _pdq.SYSTAT_TYPE_physmem_get
    if _newclass:physmem = _swig_property(_pdq.SYSTAT_TYPE_physmem_get, _pdq.SYSTAT_TYPE_physmem_set)
    __swig_setmethods__["highwater"] = _pdq.SYSTAT_TYPE_highwater_set
    __swig_getmethods__["highwater"] = _pdq.SYSTAT_TYPE_highwater_get
    if _newclass:highwater = _swig_property(_pdq.SYSTAT_TYPE_highwater_get, _pdq.SYSTAT_TYPE_highwater_set)
    __swig_setmethods__["malloc"] = _pdq.SYSTAT_TYPE_malloc_set
    __swig_getmethods__["malloc"] = _pdq.SYSTAT_TYPE_malloc_get
    if _newclass:malloc = _swig_property(_pdq.SYSTAT_TYPE_malloc_get, _pdq.SYSTAT_TYPE_malloc_set)
    __swig_setmethods__["mpl"] = _pdq.SYSTAT_TYPE_mpl_set
    __swig_getmethods__["mpl"] = _pdq.SYSTAT_TYPE_mpl_get
    if _newclass:mpl = _swig_property(_pdq.SYSTAT_TYPE_mpl_get, _pdq.SYSTAT_TYPE_mpl_set)
    __swig_setmethods__["maxN"] = _pdq.SYSTAT_TYPE_maxN_set
    __swig_getmethods__["maxN"] = _pdq.SYSTAT_TYPE_maxN_get
    if _newclass:maxN = _swig_property(_pdq.SYSTAT_TYPE_maxN_get, _pdq.SYSTAT_TYPE_maxN_set)
    __swig_setmethods__["maxTP"] = _pdq.SYSTAT_TYPE_maxTP_set
    __swig_getmethods__["maxTP"] = _pdq.SYSTAT_TYPE_maxTP_get
    if _newclass:maxTP = _swig_property(_pdq.SYSTAT_TYPE_maxTP_get, _pdq.SYSTAT_TYPE_maxTP_set)
    __swig_setmethods__["minRT"] = _pdq.SYSTAT_TYPE_minRT_set
    __swig_getmethods__["minRT"] = _pdq.SYSTAT_TYPE_minRT_get
    if _newclass:minRT = _swig_property(_pdq.SYSTAT_TYPE_minRT_get, _pdq.SYSTAT_TYPE_minRT_set)
    def __init__(self): 
        this = _pdq.new_SYSTAT_TYPE()
        try: self.this.append(this)
        except: self.this = this
    __swig_destroy__ = _pdq.delete_SYSTAT_TYPE
    __del__ = lambda self : None;
SYSTAT_TYPE_swigregister = _pdq.SYSTAT_TYPE_swigregister
SYSTAT_TYPE_swigregister(SYSTAT_TYPE)
cvar = _pdq.cvar

class TERMINAL_TYPE(_object):
    __swig_setmethods__ = {}
    __setattr__ = lambda self, name, value: _swig_setattr(self, TERMINAL_TYPE, name, value)
    __swig_getmethods__ = {}
    __getattr__ = lambda self, name: _swig_getattr(self, TERMINAL_TYPE, name)
    __repr__ = _swig_repr
    __swig_setmethods__["name"] = _pdq.TERMINAL_TYPE_name_set
    __swig_getmethods__["name"] = _pdq.TERMINAL_TYPE_name_get
    if _newclass:name = _swig_property(_pdq.TERMINAL_TYPE_name_get, _pdq.TERMINAL_TYPE_name_set)
    __swig_setmethods__["pop"] = _pdq.TERMINAL_TYPE_pop_set
    __swig_getmethods__["pop"] = _pdq.TERMINAL_TYPE_pop_get
    if _newclass:pop = _swig_property(_pdq.TERMINAL_TYPE_pop_get, _pdq.TERMINAL_TYPE_pop_set)
    __swig_setmethods__["think"] = _pdq.TERMINAL_TYPE_think_set
    __swig_getmethods__["think"] = _pdq.TERMINAL_TYPE_think_get
    if _newclass:think = _swig_property(_pdq.TERMINAL_TYPE_think_get, _pdq.TERMINAL_TYPE_think_set)
    __swig_setmethods__["sys"] = _pdq.TERMINAL_TYPE_sys_set
    __swig_getmethods__["sys"] = _pdq.TERMINAL_TYPE_sys_get
    if _newclass:sys = _swig_property(_pdq.TERMINAL_TYPE_sys_get, _pdq.TERMINAL_TYPE_sys_set)
    def __init__(self): 
        this = _pdq.new_TERMINAL_TYPE()
        try: self.this.append(this)
        except: self.this = this
    __swig_destroy__ = _pdq.delete_TERMINAL_TYPE
    __del__ = lambda self : None;
TERMINAL_TYPE_swigregister = _pdq.TERMINAL_TYPE_swigregister
TERMINAL_TYPE_swigregister(TERMINAL_TYPE)

class BATCH_TYPE(_object):
    __swig_setmethods__ = {}
    __setattr__ = lambda self, name, value: _swig_setattr(self, BATCH_TYPE, name, value)
    __swig_getmethods__ = {}
    __getattr__ = lambda self, name: _swig_getattr(self, BATCH_TYPE, name)
    __repr__ = _swig_repr
    __swig_setmethods__["name"] = _pdq.BATCH_TYPE_name_set
    __swig_getmethods__["name"] = _pdq.BATCH_TYPE_name_get
    if _newclass:name = _swig_property(_pdq.BATCH_TYPE_name_get, _pdq.BATCH_TYPE_name_set)
    __swig_setmethods__["pop"] = _pdq.BATCH_TYPE_pop_set
    __swig_getmethods__["pop"] = _pdq.BATCH_TYPE_pop_get
    if _newclass:pop = _swig_property(_pdq.BATCH_TYPE_pop_get, _pdq.BATCH_TYPE_pop_set)
    __swig_setmethods__["sys"] = _pdq.BATCH_TYPE_sys_set
    __swig_getmethods__["sys"] = _pdq.BATCH_TYPE_sys_get
    if _newclass:sys = _swig_property(_pdq.BATCH_TYPE_sys_get, _pdq.BATCH_TYPE_sys_set)
    def __init__(self): 
        this = _pdq.new_BATCH_TYPE()
        try: self.this.append(this)
        except: self.this = this
    __swig_destroy__ = _pdq.delete_BATCH_TYPE
    __del__ = lambda self : None;
BATCH_TYPE_swigregister = _pdq.BATCH_TYPE_swigregister
BATCH_TYPE_swigregister(BATCH_TYPE)

class TRANSACTION_TYPE(_object):
    __swig_setmethods__ = {}
    __setattr__ = lambda self, name, value: _swig_setattr(self, TRANSACTION_TYPE, name, value)
    __swig_getmethods__ = {}
    __getattr__ = lambda self, name: _swig_getattr(self, TRANSACTION_TYPE, name)
    __repr__ = _swig_repr
    __swig_setmethods__["name"] = _pdq.TRANSACTION_TYPE_name_set
    __swig_getmethods__["name"] = _pdq.TRANSACTION_TYPE_name_get
    if _newclass:name = _swig_property(_pdq.TRANSACTION_TYPE_name_get, _pdq.TRANSACTION_TYPE_name_set)
    __swig_setmethods__["arrival_rate"] = _pdq.TRANSACTION_TYPE_arrival_rate_set
    __swig_getmethods__["arrival_rate"] = _pdq.TRANSACTION_TYPE_arrival_rate_get
    if _newclass:arrival_rate = _swig_property(_pdq.TRANSACTION_TYPE_arrival_rate_get, _pdq.TRANSACTION_TYPE_arrival_rate_set)
    __swig_setmethods__["saturation_rate"] = _pdq.TRANSACTION_TYPE_saturation_rate_set
    __swig_getmethods__["saturation_rate"] = _pdq.TRANSACTION_TYPE_saturation_rate_get
    if _newclass:saturation_rate = _swig_property(_pdq.TRANSACTION_TYPE_saturation_rate_get, _pdq.TRANSACTION_TYPE_saturation_rate_set)
    __swig_setmethods__["sys"] = _pdq.TRANSACTION_TYPE_sys_set
    __swig_getmethods__["sys"] = _pdq.TRANSACTION_TYPE_sys_get
    if _newclass:sys = _swig_property(_pdq.TRANSACTION_TYPE_sys_get, _pdq.TRANSACTION_TYPE_sys_set)
    def __init__(self): 
        this = _pdq.new_TRANSACTION_TYPE()
        try: self.this.append(this)
        except: self.this = this
    __swig_destroy__ = _pdq.delete_TRANSACTION_TYPE
    __del__ = lambda self : None;
TRANSACTION_TYPE_swigregister = _pdq.TRANSACTION_TYPE_swigregister
TRANSACTION_TYPE_swigregister(TRANSACTION_TYPE)

class JOB_TYPE(_object):
    __swig_setmethods__ = {}
    __setattr__ = lambda self, name, value: _swig_setattr(self, JOB_TYPE, name, value)
    __swig_getmethods__ = {}
    __getattr__ = lambda self, name: _swig_getattr(self, JOB_TYPE, name)
    __repr__ = _swig_repr
    __swig_setmethods__["should_be_class"] = _pdq.JOB_TYPE_should_be_class_set
    __swig_getmethods__["should_be_class"] = _pdq.JOB_TYPE_should_be_class_get
    if _newclass:should_be_class = _swig_property(_pdq.JOB_TYPE_should_be_class_get, _pdq.JOB_TYPE_should_be_class_set)
    __swig_setmethods__["network"] = _pdq.JOB_TYPE_network_set
    __swig_getmethods__["network"] = _pdq.JOB_TYPE_network_get
    if _newclass:network = _swig_property(_pdq.JOB_TYPE_network_get, _pdq.JOB_TYPE_network_set)
    __swig_setmethods__["term"] = _pdq.JOB_TYPE_term_set
    __swig_getmethods__["term"] = _pdq.JOB_TYPE_term_get
    if _newclass:term = _swig_property(_pdq.JOB_TYPE_term_get, _pdq.JOB_TYPE_term_set)
    __swig_setmethods__["batch"] = _pdq.JOB_TYPE_batch_set
    __swig_getmethods__["batch"] = _pdq.JOB_TYPE_batch_get
    if _newclass:batch = _swig_property(_pdq.JOB_TYPE_batch_get, _pdq.JOB_TYPE_batch_set)
    __swig_setmethods__["trans"] = _pdq.JOB_TYPE_trans_set
    __swig_getmethods__["trans"] = _pdq.JOB_TYPE_trans_get
    if _newclass:trans = _swig_property(_pdq.JOB_TYPE_trans_get, _pdq.JOB_TYPE_trans_set)
    def __init__(self): 
        this = _pdq.new_JOB_TYPE()
        try: self.this.append(this)
        except: self.this = this
    __swig_destroy__ = _pdq.delete_JOB_TYPE
    __del__ = lambda self : None;
JOB_TYPE_swigregister = _pdq.JOB_TYPE_swigregister
JOB_TYPE_swigregister(JOB_TYPE)

class NODE_TYPE(_object):
    __swig_setmethods__ = {}
    __setattr__ = lambda self, name, value: _swig_setattr(self, NODE_TYPE, name, value)
    __swig_getmethods__ = {}
    __getattr__ = lambda self, name: _swig_getattr(self, NODE_TYPE, name)
    __repr__ = _swig_repr
    __swig_setmethods__["devtype"] = _pdq.NODE_TYPE_devtype_set
    __swig_getmethods__["devtype"] = _pdq.NODE_TYPE_devtype_get
    if _newclass:devtype = _swig_property(_pdq.NODE_TYPE_devtype_get, _pdq.NODE_TYPE_devtype_set)
    __swig_setmethods__["sched"] = _pdq.NODE_TYPE_sched_set
    __swig_getmethods__["sched"] = _pdq.NODE_TYPE_sched_get
    if _newclass:sched = _swig_property(_pdq.NODE_TYPE_sched_get, _pdq.NODE_TYPE_sched_set)
    __swig_setmethods__["devname"] = _pdq.NODE_TYPE_devname_set
    __swig_getmethods__["devname"] = _pdq.NODE_TYPE_devname_get
    if _newclass:devname = _swig_property(_pdq.NODE_TYPE_devname_get, _pdq.NODE_TYPE_devname_set)
    __swig_setmethods__["visits"] = _pdq.NODE_TYPE_visits_set
    __swig_getmethods__["visits"] = _pdq.NODE_TYPE_visits_get
    if _newclass:visits = _swig_property(_pdq.NODE_TYPE_visits_get, _pdq.NODE_TYPE_visits_set)
    __swig_setmethods__["service"] = _pdq.NODE_TYPE_service_set
    __swig_getmethods__["service"] = _pdq.NODE_TYPE_service_get
    if _newclass:service = _swig_property(_pdq.NODE_TYPE_service_get, _pdq.NODE_TYPE_service_set)
    __swig_setmethods__["demand"] = _pdq.NODE_TYPE_demand_set
    __swig_getmethods__["demand"] = _pdq.NODE_TYPE_demand_get
    if _newclass:demand = _swig_property(_pdq.NODE_TYPE_demand_get, _pdq.NODE_TYPE_demand_set)
    __swig_setmethods__["resit"] = _pdq.NODE_TYPE_resit_set
    __swig_getmethods__["resit"] = _pdq.NODE_TYPE_resit_get
    if _newclass:resit = _swig_property(_pdq.NODE_TYPE_resit_get, _pdq.NODE_TYPE_resit_set)
    __swig_setmethods__["utiliz"] = _pdq.NODE_TYPE_utiliz_set
    __swig_getmethods__["utiliz"] = _pdq.NODE_TYPE_utiliz_get
    if _newclass:utiliz = _swig_property(_pdq.NODE_TYPE_utiliz_get, _pdq.NODE_TYPE_utiliz_set)
    __swig_setmethods__["qsize"] = _pdq.NODE_TYPE_qsize_set
    __swig_getmethods__["qsize"] = _pdq.NODE_TYPE_qsize_get
    if _newclass:qsize = _swig_property(_pdq.NODE_TYPE_qsize_get, _pdq.NODE_TYPE_qsize_set)
    __swig_setmethods__["avqsize"] = _pdq.NODE_TYPE_avqsize_set
    __swig_getmethods__["avqsize"] = _pdq.NODE_TYPE_avqsize_get
    if _newclass:avqsize = _swig_property(_pdq.NODE_TYPE_avqsize_get, _pdq.NODE_TYPE_avqsize_set)
    def __init__(self): 
        this = _pdq.new_NODE_TYPE()
        try: self.this.append(this)
        except: self.this = this
    __swig_destroy__ = _pdq.delete_NODE_TYPE
    __del__ = lambda self : None;
NODE_TYPE_swigregister = _pdq.NODE_TYPE_swigregister
NODE_TYPE_swigregister(NODE_TYPE)


def CreateClosed(*args):
  return _pdq.CreateClosed(*args)
CreateClosed = _pdq.CreateClosed

def CreateClosed_p(*args):
  return _pdq.CreateClosed_p(*args)
CreateClosed_p = _pdq.CreateClosed_p

def CreateOpen(*args):
  return _pdq.CreateOpen(*args)
CreateOpen = _pdq.CreateOpen

def CreateOpen_p(*args):
  return _pdq.CreateOpen_p(*args)
CreateOpen_p = _pdq.CreateOpen_p

def CreateNode(*args):
  return _pdq.CreateNode(*args)
CreateNode = _pdq.CreateNode

def CreateMultiNode(*args):
  return _pdq.CreateMultiNode(*args)
CreateMultiNode = _pdq.CreateMultiNode

def GetStreamsCount():
  return _pdq.GetStreamsCount()
GetStreamsCount = _pdq.GetStreamsCount

def GetNodesCount():
  return _pdq.GetNodesCount()
GetNodesCount = _pdq.GetNodesCount

def GetResponse(*args):
  return _pdq.GetResponse(*args)
GetResponse = _pdq.GetResponse

def PDQ_GetResidenceTime(*args):
  return _pdq.PDQ_GetResidenceTime(*args)
PDQ_GetResidenceTime = _pdq.PDQ_GetResidenceTime

def GetThruput(*args):
  return _pdq.GetThruput(*args)
GetThruput = _pdq.GetThruput

def PDQ_GetLoadOpt(*args):
  return _pdq.PDQ_GetLoadOpt(*args)
PDQ_GetLoadOpt = _pdq.PDQ_GetLoadOpt

def GetUtilization(*args):
  return _pdq.GetUtilization(*args)
GetUtilization = _pdq.GetUtilization

def GetQueueLength(*args):
  return _pdq.GetQueueLength(*args)
GetQueueLength = _pdq.GetQueueLength

def PDQ_GetThruMax(*args):
  return _pdq.PDQ_GetThruMax(*args)
PDQ_GetThruMax = _pdq.PDQ_GetThruMax

def Init(*args):
  return _pdq.Init(*args)
Init = _pdq.Init

def Report():
  return _pdq.Report()
Report = _pdq.Report

def SetDebug(*args):
  return _pdq.SetDebug(*args)
SetDebug = _pdq.SetDebug

def SetDemand(*args):
  return _pdq.SetDemand(*args)
SetDemand = _pdq.SetDemand

def SetDemand_p(*args):
  return _pdq.SetDemand_p(*args)
SetDemand_p = _pdq.SetDemand_p

def SetVisits(*args):
  return _pdq.SetVisits(*args)
SetVisits = _pdq.SetVisits

def SetVisits_p(*args):
  return _pdq.SetVisits_p(*args)
SetVisits_p = _pdq.SetVisits_p

def Solve(*args):
  return _pdq.Solve(*args)
Solve = _pdq.Solve

def SetWUnit(*args):
  return _pdq.SetWUnit(*args)
SetWUnit = _pdq.SetWUnit

def SetTUnit(*args):
  return _pdq.SetTUnit(*args)
SetTUnit = _pdq.SetTUnit

def SetComment(*args):
  return _pdq.SetComment(*args)
SetComment = _pdq.SetComment

def GetComment():
  return _pdq.GetComment()
GetComment = _pdq.GetComment

def PrintNodes():
  return _pdq.PrintNodes()
PrintNodes = _pdq.PrintNodes

def GetNode(*args):
  return _pdq.GetNode(*args)
GetNode = _pdq.GetNode

def getjob(*args):
  return _pdq.getjob(*args)
getjob = _pdq.getjob


